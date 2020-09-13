package ovm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceOvmVM(t *testing.T) {
	resourceName := "data.ovm_vm.ovm_template"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreChecks(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOvmVMDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOvmVMDataSourceID(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "oracle77-devans-dev.0"),
				),
			},
		},
	})
}

func testAccCheckOvmVMDataSourceID(n string) resource.TestCheckFunc {
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

const testAccCheckOvmVMDataSourceConfig = `
data "ovm_repository" "repo" {
	name = "ovm-corp-repository"
}

data "ovm_vm" "ovm_template" {
	name         = "oracle77-devans-dev.0"
	repositoryid = data.ovm_repository.repo
}
`
