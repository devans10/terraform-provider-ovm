# Data Source ServerPool

Data source to return the ID of a ServerPool object

## Example Usage
```sh
data "ovm_serverpool" "pool" {
    name = "my-serverpool"
}
```

## Attributes

The following arguments are supported:

+ `id` - ID of the ServerPool
+ `value` - Same as ID
+ `uri` - URI of the ServerPool object
+ `name` - Name of the ServerPool
+ `type` - OVM object model