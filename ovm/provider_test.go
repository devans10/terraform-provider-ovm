package ovm

import (
	"os"
	"testing"

	//"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	//"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ovm": testAccProvider,
	}
}

// func TestMain(m *testing.M) {
// 	acctest.UseBinaryDriver("ovm", Provider)
// 	resource.TestMain(m)
// }

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreChecks(t *testing.T) {

	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC to run terraform acceptance tests (provider connection is required)")
	}

	entrypoint := os.Getenv("OVM_ENDPOINT")
	username := os.Getenv("OVM_USERNAME")
	password := os.Getenv("OVM_PASSWORD")
	if entrypoint == "" {
		t.Fatalf("OVM_ENDPOINT must be set for acceptance tests")
	}
	if (username != "") && (password == "") {
		t.Fatalf("OVM_PASSWORD must be set if OVM_USERNAME is set for acceptance tests")
	}
}

func testAccProviderMeta(t *testing.T) (interface{}, error) {
	t.Helper()
	d := schema.TestResourceDataRaw(t, testAccProvider.Schema, make(map[string]interface{}))
	return providerConfigure(d)
}
