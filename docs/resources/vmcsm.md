# Resource VMCSM

Resource to create a VM Clone Storage Mapping

## Example

```sh
data "ovm_repository" "repo" {
  name = var.repository
}

data "ovm_vm" "ovm_template" {
  name         = var.machine_image
  repositoryid = data.ovm_repository.repo
}

data "ovm_serverpool" "serverpool" {
  name = var.server_pool
}

data "ovm_network" "network" {
  name = var.network
}

# //Creating VmCloneCustomizer
resource "ovm_vmcd" "oel7_tmpl_cst" {
  vmid        = data.ovm_vm.ovm_template.id
  name        = "oe7_tmpl_cst"
  description = "Desc oel7 cust"
}

# //Defining Vm Clone Storage Mapping
resource "ovm_vmcsm" "oel7_vmclonestoragemapping" {
  for_each = toset(data.ovm_vm.ovm_template.vmdiskmappingids.*.value)

  vmdiskmappingid     = each.key
  vmclonedefinitionid = ovm_vmcd.oel7_tmpl_cst.id
  repositoryid        = data.ovm_repository.repo
  name                = "oel_cust_storage"
  clonetype           = "SPARSE_COPY"
}
```

## Attributes

The following arguments are supported:

+ `vmdiskmappingid`
+ `vmclonedefinitionid` 
+ `repositoryid`
+ `clonetype`
+ `name`