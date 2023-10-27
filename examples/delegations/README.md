This example shows the use of delegations on subnets. Delegations permit specific azure services to operate within a subnet, essentially granting them tailored permissions.

This helps in scenarios where certain services need specialized access to function correctly.

## Usage

```hcl
module "network" {
  source = "github.com/cloudnationhq/az-cn-module-tf-vnet"

  naming = local.naming

  vnet = {
    name          = module.naming.virtual_network.name
    location      = module.rg.groups.demo.location
    resourcegroup = module.rg.groups.demo.name
    cidr          = ["10.18.0.0/16"]

    subnets = {
      sn1 = {
        cidr = ["10.18.1.0/24"]
        delegations = {
          sql = {
            name = "Microsoft.Sql/managedInstances"
            actions = [
              "Microsoft.Network/virtualNetworks/subnets/join/action",
              "Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action",
              "Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action",
            ]
          }
        }
      }
      sn2 = {
        cidr = ["10.18.2.0/24"]
        delegations = {
          web = { name = "Microsoft.Web/serverFarms" }
        }
      }
    }
  }
}
```

## Types

```hcl
variable "vnet" {
  type = object({
    name          = string
    location      = string
    resourcegroup = string
    cidr          = list(string)

    subnets = optional(map(object({
      cidr       = list(string)
      delegations = optional(map(object({
        name    = string
        actions = optional(list(string), null)
      })), null)
    })), null)
  })
}

variable "naming" {
  type = map(string)
}
```

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.61 |
| <a name="requirement_random"></a> [random](#requirement\_random) | ~> 3.5.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | ~> 3.61 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_naming"></a> [naming](#input\_naming) | n/a | `map(string)` | n/a | yes |
| <a name="input_vnet"></a> [vnet](#input\_vnet) | n/a | `any` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_subnets"></a> [subnets](#output\_subnets) | n/a |
| <a name="output_subscriptionId"></a> [subscriptionId](#output\_subscriptionId) | n/a |
| <a name="output_vnet"></a> [vnet](#output\_vnet) | n/a |
