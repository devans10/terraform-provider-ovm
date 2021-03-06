package main

import (
	"github.com/devans10/terraform-provider-ovm/ovm"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ovm.Provider,
	})
}
