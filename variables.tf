//variable "vnet" {
//  type = any
//}


variable "vnet" {
  type = object({
    name : string
    resourcegroup : string
    location : string
    cidr : list(string)
    subnets : optional(map(object({
      name : optional(string)
      cidr : list(string)
      endpoints : optional(list(string))
      enforce_priv_link_service : optional(bool)
      enforce_priv_link_endpoint : optional(bool)
      delegations : optional(map(object({
        name : string
        actions : optional(list(string))
      })))
      rules : optional(map(object({
        name : string
        priority : number
        direction : string
        access : string
        protocol : string
        description : optional(string)
        source_port_range : optional(string)
        source_port_ranges : optional(list(string))
        destination_port_range : optional(string)
        destination_port_ranges : optional(list(string))
        source_address_prefix : optional(string)
        source_address_prefixes : optional(list(string))
        destination_address_prefix : optional(string)
        destination_address_prefixes : optional(list(string))
      })))
      nsg : optional(object({
        name : optional(string)
      }))
      route : optional(object({
        name : optional(string)
        routes : optional(map(object({
          address_prefix : optional(string)
          next_hop_type : optional(string)
          next_hop_in_ip_address : optional(string)
        })))
      }))
      route_table : optional(string)
      shd_route_table : optional(string)
    })))
  })
}

variable "naming" {
  type = map(string)
}
