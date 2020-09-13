package ovm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceOvmRepository(t *testing.T) {
	resourceName := "data.ovm_repository.repo"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreChecks(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOvmRepositoryDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOvmRepositoryDataSourceID(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "ovm-corp-repository"),
					resource.TestCheckResourceAttr(resourceName, "type", "com.oracle.ovm.mgr.ws.model.Repository"),
				),
			},
		},
	})
}

func testAccCheckOvmRepositoryDataSourceID(n string) resource.TestCheckFunc {
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

const testAccCheckOvmRepositoryDataSourceConfig = `
data "ovm_repository" "repo" {
	name = "ovm-corp-repository"
}
`
