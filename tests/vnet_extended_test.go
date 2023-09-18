package main

import (
	"context"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/cloudnationhq/az-cn-module-tf-vnet/shared"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

type VnetDetails struct {
	ResourceGroupName string
	Name              string
	Subnets           []SubnetDetails
}

type SubnetDetails struct {
	Name                   string
	NetworkSecurityGroupID string
}

type ClientSetup struct {
	SubscriptionID     string
	VirtualNetworkClient *armnetwork.VirtualNetworksClient
	SubnetsClient         *armnetwork.SubnetsClient
}

func (details *VnetDetails) GetVnet(t *testing.T, client *armnetwork.VirtualNetworksClient) *armnetwork.VirtualNetwork {
	resp, err := client.Get(context.Background(), details.ResourceGroupName, details.Name, nil)
	require.NoError(t, err, "Failed to get vnet details")
	return &resp.VirtualNetwork
}

func (details *VnetDetails) GetSubnets(t *testing.T, client *armnetwork.SubnetsClient) []SubnetDetails {
	pager := client.NewListPager(details.ResourceGroupName, details.Name, nil)

	var subnets []SubnetDetails
	for {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err, "Failed to list subnets")
		for _, subnet := range page.Value {
			subnets = append(subnets, SubnetDetails{
				Name:                   *subnet.Name,
				NetworkSecurityGroupID: *subnet.Properties.NetworkSecurityGroup.ID,
			})
		}

		if page.NextLink == nil || len(*page.NextLink) == 0 {
			break
		}
	}
	return subnets
}

func (setup *ClientSetup) InitializeVirtualNetworkClient(t *testing.T, cred *azidentity.DefaultAzureCredential) {
	var err error
	setup.VirtualNetworkClient, err = armnetwork.NewVirtualNetworksClient(setup.SubscriptionID, cred, nil)
	require.NoError(t, err, "Failed to initialize virtual network client")
}

func (setup *ClientSetup) InitializeSubnetsClient(t *testing.T, cred *azidentity.DefaultAzureCredential) {
	var err error
	setup.SubnetsClient, err = armnetwork.NewSubnetsClient(setup.SubscriptionID, cred, nil)
	require.NoError(t, err, "Failed to initialize subnets client")
}


func TestVirtualNetwork(t *testing.T) {
	t.Run("VerifyVnet", func(t *testing.T) {
		t.Parallel()

		cred, err := azidentity.NewDefaultAzureCredential(nil)
		require.NoError(t, err, "Failed to get credentials")

		tfOpts := shared.GetTerraformOptions("../examples/complete")
		defer shared.Cleanup(t, tfOpts)
		terraform.InitAndApply(t, tfOpts)

		vnetMap := terraform.OutputMap(t, tfOpts, "vnet")
		subscriptionID := terraform.Output(t, tfOpts, "subscriptionId")

		vnetDetails := &VnetDetails{
			ResourceGroupName: vnetMap["resource_group_name"],
			Name:              vnetMap["name"],
		}

		clientSetup := &ClientSetup{SubscriptionID: subscriptionID}
		clientSetup.InitializeVirtualNetworkClient(t, cred)
		clientSetup.InitializeSubnetsClient(t, cred)

		vnet := vnetDetails.GetVnet(t, clientSetup.VirtualNetworkClient)
		vnetDetails.Subnets = vnetDetails.GetSubnets(t, clientSetup.SubnetsClient)

		t.Run("VerifyVnetDetails", func(t *testing.T) {
			verifyVnetDetails(t, vnetDetails, vnet)
		})

		//t.Run("VerifySubnetsExist", func(t *testing.T) {
		//	verifySubnetsExist(t, vnetDetails)
		//})
	})
}

func verifyVnetDetails(t *testing.T, details *VnetDetails, vnet *armnetwork.VirtualNetwork) {
	t.Helper()

	require.Equal(
		t,
		details.Name,
		*vnet.Name,
		"Vnet name does not match expected value",
	)

	require.Equal(
		t,
		"Succeeded",
		string(*vnet.Properties.ProvisioningState),
		"Vnet provisioning state is not 'Succeeded'",
	)

	require.True(
		t,
		strings.HasPrefix(details.Name, "vnet"),
		"Vnet name does not begin with the right abbreviation",
	)
}

//func verifySubnetsExist(t *testing.T, details *VnetDetails) {
//	t.Helper()

//	for _, subnet := range details.Subnets {
//		require.NotEmpty(
//			t,
//			subnet.Name,
//			"Subnet name is empty",
//		)

//		require.NotEmpty(
//			t,
//			subnet.NetworkSecurityGroupID,
//			"No network security group association found",
//		)

//		require.True(
//			t,
//			strings.HasPrefix(subnet.Name, "snet"),
//			"Subnet name does not begin with the right abbreviation",
//		)
//	}
//}
