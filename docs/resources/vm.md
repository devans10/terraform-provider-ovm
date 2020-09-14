# Resource VM

Resource to create a VM

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

resource "ovm_vm" "cloneoel7" {
  name = "cloneoel7Vm"

  repositoryid = data.ovm_repository.repo
  serverpoolid = data.ovm_serverpool.serverpool

  memorylimit         = 4096
  memory              = 4096
  cpucount            = 4
  cpucountlimit       = 4
  vmdomaintype        = "XEN_HVM_PV_DRIVERS"
  imageid             = data.ovm_vm.ovm_template.id
  ostype              = "Oracle Linux 7"
  vmclonedefinitionid = ovm_vmcd.oel7_tmpl_cst.id

  virtualnic {
    networkid = data.ovm_network.network
  }

  sendmessages = {
    "com.oracle.linux.network.hostname"    = "cloneoel7vm"
    "com.oracle.linux.network.device.0"    = "eth0"
    "com.oracle.linux.network.bootproto.0" = "dhcp"
    "com.oracle.linux.network.onboot.0"    = "yes"
    "com.oracle.linux.root-password"       = "Welcome!"
  }

  depends_on = [ovm_vmcsm.oel7_vmclonestoragemapping]
}
```

## Attributes

The following arguments are supported:

+ `name`
+ `repositoryid` 
+ `serverpoolid`
+ `affinitygroupids`
+ `architecture`
+ `bootorder`
+ `cpucount` 
+ `cpucountlimit`
+ `cpupriority`
+ `cpuutilizationcap`
+ `description`
+ `disklimit`
+ `generation`
+ `guestdriverversion` 
+ `highavailability`
+ `hugepagesenabled`
+ `kernelversion`
+ `keymapname`
+ `locked`
+ `memory`
+ `memorylimit`
+ `networkinstallpath`
+ `origin`
+ `ostype`
+ `osversion`
+ `pinnedcpus`
+ `readonly`
+ `resourcegroupids`
+ `restartactiononcrash`
+ `restartactiononpoweroff`
+ `restartactiononrestart`
+ `serverid`
+ `sslvncport`
+ `sslttyport`
+ `userdata`
+ `virtualnicids`
+ `vmapiversion`
+ `vmclonedefinitionids`
+ `vmconfigfileabsolutepath`
+ `vmconfigfilemountedpath`
+ `vmdiskmappingids`
+ `vmdomaintype`
+ `vmmousetype`
+ `vmrunstate`
+ `vmstartpolicy`
+ `vmclonedefinitionid`
+ `imageid`
+ `virtualnic`
+ `sendmessages`


