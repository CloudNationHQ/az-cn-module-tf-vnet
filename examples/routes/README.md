This example underscores the implementation of routes within subnets. Routes direct traffic flow, ensuring optimized and secure network navigation. This is crucial in scenarios where specific traffic patterns or destinations need to be enforced for a subnet's resources.

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
        routes = {
          rt1 = {
            address_prefix = "Storage"
            next_hop_type  = "Internet"
          }
        }
      }
    }
  }
}
```

In situations where several subnets should share the same route table, the following configuration can be employed:

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
        cidr        = ["10.18.1.0/24"]
        route_table = "shd"
      },
      sn2 = {
        cidr        = ["10.18.2.0/24"]
        route_table = "shd"
      }
    }

    route_tables = {
      shd = {
        routes = {
          rt1 = { address_prefix = "0.0.0.0/0", next_hop_type = "Internet" }
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
    name          : string
    location      : string
    resourcegroup : string
    cidr          : list(string)
    subnets       : map(object({
      cidr        : list(string)
      route_table : optional(string)
      routes      : optional(map(object({
        address_prefix : string
        next_hop_type  : string
      })))
    }))
    route_tables   : optional(map(object({
      routes      : map(object({
        address_prefix : string
        next_hop_type  : string
      }))
    })))
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

## Resources

| Name | Type |
| :-- | :-- |
| [azurerm_network_security_group.nsg](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_security_group) | resource |
| [azurerm_route_table.rt](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/route_table) | resource |
| [azurerm_route_table.shd_rt](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/route_table) | resource |
| [azurerm_subnet.subnets](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet) | resource |
| [azurerm_subnet_network_security_group_association.nsg_as](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet_network_security_group_association) | resource |
| [azurerm_subnet_route_table_association.rt_as](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/subnet_route_table_association) | resource |
| [azurerm_virtual_network.vnet](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_network) | resource |
| [azurerm_virtual_network_dns_servers.dns](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_network_dns_servers) | resource |
| [azurerm_subscription.current](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/data-sources/subscription) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_naming"></a> [naming](#input\_naming) | contains naming convention | `map(string)` | n/a | yes |
| <a name="input_vnet"></a> [vnet](#input\_vnet) | describes vnet related configuration | `any` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_subnets"></a> [subnets](#output\_subnets) | subnet configuration specifics |
| <a name="output_subscriptionId"></a> [subscriptionId](#output\_subscriptionId) | contains the current subscriptionId |
| <a name="output_vnet"></a> [vnet](#output\_vnet) | vnet setup details |

## Testing

make test TF_PATH=routes
