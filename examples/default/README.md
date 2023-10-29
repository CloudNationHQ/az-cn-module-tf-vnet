This example illustrates the default virtual network setup, in its simplest form. It lays out a foundational network structure, guaranteeing protected and contained communication within its scope.

## Usage: simple

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
      demo = {
        cidr = ["10.18.3.0/24"]
        endpoints = [
          "Microsoft.Storage",
          "Microsoft.Sql"
        ]
      }
    }
  }
}
```

## Usage: multiple

Additionally, for certain scenarios, the example below highlights the ability to use multiple virtual networks, enabling a broader network setup.

```hcl
module "network" {
  source = "../../"

  for_each = local.vnets

  naming = local.naming
  vnet   = each.value
}
```

The module iterates over ```local.vnets```, creating a virtual network for each entry.

```hcl
locals {
  vnets = {
    vnet1 = {
      name          = join("-", [module.naming.virtual_network.name, "001"])
      location      = module.rg.groups.demo.location
      resourcegroup = module.rg.groups.demo.name
      cidr          = ["10.18.0.0/16"]

      subnets = {
        sql = {
          cidr = ["10.18.1.0/24"]
          endpoints = [
            "Microsoft.Sql"
          ]
        },
        ws = {
          cidr = ["10.18.2.0/24"]
          delegations = {
            databricks = {
              name = "Microsoft.Databricks/workspaces"
            }
          }
        }
      }
    },
    vnet2 = {
      name          = join("-", [module.naming.virtual_network.name, "002"])
      location      = module.rg.groups.demo.location
      resourcegroup = module.rg.groups.demo.name
      cidr          = ["10.20.0.0/16"]

      subnets = {
        plink = {
          cidr = ["10.20.1.0/24"]
          endpoints = [
            "Microsoft.Storage",
          ]
        }
      }
    }
  }
}
```
