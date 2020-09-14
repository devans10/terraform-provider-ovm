# Data Source Network

Data source to return the ID of a Network object

## Example Usage
```sh
data "ovm_network" "network" {
    name = "vm-public"
}
```

## Attributes

The following arguments are supported:

+ `id` - ID of the Network
+ `value` - Same as ID
+ `uri` - URI of the Network object
+ `name` - Name of the Network
+ `type` - OVM object model