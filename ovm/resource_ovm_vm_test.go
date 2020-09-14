package ovm

import (
	"fmt"
	"testing"
	"time"

	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceOvmVM(t *testing.T) {
	resourceName := "ovm_vm.cloneoel7"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreChecks(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOvmVMDestroy,
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

	time.Sleep(5 * time.Second)

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
# //Creating VmCloneCustomizer
resource "ovm_vmcd" "oel7_tmpl_cst" {
	vmid        = "0004fb00000600005d472b3bdc6b4cb2"
	name        = "oe7_tmpl_cst"
	description = "Desc oel7 cust"
}
  
# //Defining Vm Clone Storage Mapping
resource "ovm_vmcsm" "oel7_vmclonestoragemapping" {
	vmdiskmappingid     = "0004fb00001300001356f01b2ed0d6f3"
	vmclonedefinitionid = ovm_vmcd.oel7_tmpl_cst.id
	repositoryid        = {
		"id"    = "0004fb000003000019686ad4c15c1306"
		"name"  = "ovm-corp-repository"
		"type"  = "com.oracle.ovm.mgr.ws.model.Repository"
		"uri"   = "https://localhost:7002//ovm/core/wsapi/rest/Repository/0004fb000003000019686ad4c15c1306"
		"value" = "0004fb000003000019686ad4c15c1306"
	}
	name                = "oel_cust_storage"
	clonetype           = "SPARSE_COPY"
}
  
resource "ovm_vm" "cloneoel7" {
	name = "cloneoel7Vm"
  
	repositoryid = {
		"id"    = "0004fb000003000019686ad4c15c1306"
		"name"  = "ovm-corp-repository"
		"type"  = "com.oracle.ovm.mgr.ws.model.Repository"
		"uri"   = "https://localhost:7002//ovm/core/wsapi/rest/Repository/0004fb000003000019686ad4c15c1306"
		"value" = "0004fb000003000019686ad4c15c1306"
	}
	serverpoolid = {
		"id"    = "0004fb00000200008f1fe276aab59de8"
		"name"  = "xen-pool2"
		"type"  = "com.oracle.ovm.mgr.ws.model.ServerPool"
		"uri"   = "https://localhost:7002//ovm/core/wsapi/rest/ServerPool/0004fb00000200008f1fe276aab59de8"
		"value" = "0004fb00000200008f1fe276aab59de8"
	}
	memorylimit         = 4096
	memory              = 4096
	cpucount            = 4
	cpucountlimit       = 4
	vmdomaintype        = "XEN_HVM_PV_DRIVERS"
	imageid             = "0004fb00000600005d472b3bdc6b4cb2"
	ostype              = "Oracle Linux 7"
	vmclonedefinitionid = ovm_vmcd.oel7_tmpl_cst.id
  
	virtualnic {
		networkid  = {
			"id"    = "1063b8b5eb"
			"name"  = "vm-corp-public"
			"type"  = "com.oracle.ovm.mgr.ws.model.Network"
			"uri"   = "https://localhost:7002//ovm/core/wsapi/rest/Network/1063b8b5eb"
			"value" = "1063b8b5eb"
		}
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
`
