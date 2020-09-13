package ovm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceOvmServerPool(t *testing.T) {
	resourceName := "data.ovm_serverpool.serverpool"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreChecks(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOvmServerPoolDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "xen-pool2"),
				),
			},
		},
	})
}

func testAccCheckOvmServerPoolDataSourceID(n string) resource.TestCheckFunc {
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

const testAccCheckOvmServerPoolDataSourceConfig = `
data "ovm_serverpool" "serverpool" {
	name = "xen-pool2"
}
`
