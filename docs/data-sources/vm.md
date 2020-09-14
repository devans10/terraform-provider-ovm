# Data Source Network

Data source to return the ID of a Network object

## Example Usage
```sh
data "ovm_vm" "virtualmachine" {
    name = "my-tmpl-vm"
}
```

## Attributes

The following arguments are supported:

+ `name`
+ `repositoryid` 
+ `cpucount` 
+ `cpucountlimit` 
+ `highavailability`
+ `hugepagesenabled`
+ `memory`
+ `memorylimit`
+ `ostype`
+ `vmdomaintype`
+ `vmmousetype`
+ `osversion`
+ `vmdiskmappingids`
+ `virtualnicids`
+ `serverpoolid`