## Usage: virtual hub connection

```hcl
module "vhub-connection" {
  source = "github.com/cloudnationhq/az-cn-module-tf-vnet/vhub-connection"

  providers = {
    azurerm = azurerm.connectivity
  }

  virtual_hub = {
    name          = "vhub-westeurope"
    resourcegroup = "rg-vwan-shared"
    connection    = module.naming.virtual_hub_connection.name
    vnet          = module.network.vnet.id
  }
}
```
