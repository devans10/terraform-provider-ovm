package ovm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceOvmNetwork(t *testing.T) {
	resourceName := "data.ovm_network.network"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreChecks(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOvmNetworkDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOvmNetworkDataSourceID(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "vm-corp-public"),
					resource.TestCheckResourceAttr(resourceName, "type", "com.oracle.ovm.mgr.ws.model.Network"),
					resource.TestCheckResourceAttr(resourceName, "value", "1063b8b5eb"),
				),
			},
		},
	})
}

func testAccCheckOvmNetworkDataSourceID(n string) resource.TestCheckFunc {
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

const testAccCheckOvmNetworkDataSourceConfig = `
data "ovm_network" "network" {
	name = "vm-corp-public"
}
`
