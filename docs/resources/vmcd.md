# Resource VMCD

Resource to create a VM Clone Customizer

## Example

```sh
data "ovm_repository" "repo" {
  name = var.repository
}

data "ovm_vm" "ovm_template" {
  name         = var.machine_image
  repositoryid = data.ovm_repository.repo
}

# //Creating VmCloneCustomizer
resource "ovm_vmcd" "oel7_tmpl_cst" {
  vmid        = data.ovm_vm.ovm_template.id
  name        = "oe7_tmpl_cst"
  description = "Desc oel7 cust"
}
```

## Attributes

The following arguments are supported:

+ `vmid`
+ `name` 
+ `description`
