# terraform-provider-ovm #

### Dependencies ###

This project is a [terraform](http://www.terraform.io/) provider for [OVM](https://www.oracle.com/virtualization/oracle-vm-server-for-x86/index.html)

You should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

This relies on the [go-ovm-helper](https://github.com/devans10/go-ovm-helper) library. To
get that: `go get github.com/devans10/go-ovm-helper/ovmhelper`.

You'll also need the libraries from terraform.  Check out those docs under [plugin basics](http://www.terraform.io/docs/plugins/basics.html)

### Build ###

Run `go install github.com/devans10/terraform-provider-ovm`

### Install ###

Add the following to your terraform module

```
terraform {
  required_providers {
    ovm = {
      source = "devans10/ovm"
      version = "1.0.0"
    }
  }
}
```

## Usage ##

**Configure the Provider**

***Configure in TF configuration***

```
provider "ovm" {
  user       = "${var.ovm_username}"
  password   = "${var.ovm_password}"
  entrypoint = "${var.ovm_endpoint}"
}
```

***Configure in environment***

Set username(`OVM_USERNAME`) and password(`OVM_PASSWORD`) and endpoint(`OVM_ENDPOINT`) in environment
```
provider "ovm" {}
```

**Basic vm provision**

Create one vm and create two virtual disks and mapp them to the vm
```
resource "ovm_vm" "vm1" {
  name          = "vm1"
  repositoryid  = "${var.vm_repositoryid}"
  serverpoolid  = "${var.serverpoolid}"
  vmdomaintype  = "XEN_HVM"
  cpucount      = 2
  cpucountlimit = 2
  memory        = 512 //MB
}

resource "ovm_vd" "vm1_virtualdisk" {
  count        = 2
  name         = "vm1_vd${count.index}"
  sparse       = true
  shareable    = false
  repositoryid = "${var.vd_repositoryid}"
  size         = 104857600 //bytes
}

resource "ovm_vdm" "vm1_vdm" {
  count       = 2
  vmid        = "${ovm_vm.vm1.id}"
  vdid        = "${element(ovm_vd.vm1.*.id, count.index)}"
  name        = "vm1_vdm_2${count.index}"
  slot        = "${count.index}"
  description = "Virtual disk mapping for vm1 and vm1_vdm_2${count.index}"
}
```

**Create VM from a Template**

```
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
