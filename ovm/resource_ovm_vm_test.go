package ovm

import (
	"fmt"
	"testing"

	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceOvmVM(t *testing.T) {
	resourceName := "ovm_vm.cloneoel7"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreChecks(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOvmVMConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOvmVMID(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "cloneoel7Vm"),
					resource.TestCheckResourceAttr(resourceName, "memorylimit", "4096"),
					resource.TestCheckResourceAttr(resourceName, "memory", "4096"),
					resource.TestCheckResourceAttr(resourceName, "cpucountlimit", "4"),
					resource.TestCheckResourceAttr(resourceName, "cpucount", "4"),
					resource.TestCheckResourceAttr(resourceName, "vmdomaintype", "XEN_HVM_PV_DRIVERS"),
				),
			},
		},
	})
}

func testAccCheckOvmVMID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find Repository data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Repository data source ID not set")
		}
		return nil
	}
}

func testAccCheckOvmVMDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ovmhelper.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ovm_vm" {
			continue
		}

		_, err := client.Vms.Read(rs.Primary.ID)
		if err != nil {
			return nil
		}
		return fmt.Errorf("vm '%s' stil exists", rs.Primary.ID)
	}

	return nil
}

const testAccCheckOvmVMConfig = `
data "ovm_repository" "repo" {
	name = "ovm-corp-repository"
}
  
data "ovm_vm" "ovm_template" {
	name         = "oracle77-devans-dev.0"
	repositoryid = data.ovm_repository.repo
}
  
data "ovm_serverpool" "serverpool" {
	name = "xen-pool2"
}
  
# //Creating VmCloneCustomizer
resource "ovm_vmcd" "oel7_tmpl_cst" {
	vmid        = data.ovm_vm.ovm_template.id
	name        = "oe7_tmpl_cst"
	description = "Desc oel7 cust"
}
  
# //Defining Vm Clone Storage Mapping
resource "ovm_vmcsm" "oel7_vmclonestoragemapping" {
	vmdiskmappingid     = "0004fb00001300001356f01b2ed0d6f3"
	vmclonedefinitionid = ovm_vmcd.oel7_tmpl_cst.id
	repositoryid        = data.ovm_repository.repo
	name                = "oel_cust_storage"
	clonetype           = "SPARSE_COPY"
}
  
  # //Defining Vm Clone Network Mappings.
  # # resource "ovm_vmcnm" "oel7_vmclonenetworkmapping" {
  # #   networkid           = "	1063b8b5eb"
  # #   vmclonedefinitionid = "${ovm_vmcd.oe7_tmpl_cst.id}"
  # #   virtualnicid        = "${var.virtualnicid}"
  # #   name                = "oel_cust_network"
  # # }
  
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
  
	sendmessages = {
	  "com.oracle.linux.network.hostname"    = "cloneoel7vm"
	  "com.oracle.linux.network.device.0"    = "eth0"
	  "com.oracle.linux.network.bootproto.0" = "dhcp"
	  "com.oracle.linux.network.onboot.0"    = "yes"
	  "com.oracle.linux.root-password"       = "Welcome!"
	}
  
	depends_on = [ovm_vmcsm.oel7_vmclonestoragemapping]
}
`
