package ovm

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mitchellh/mapstructure"
)

// Provider - sets the provider settings to connect to OVM
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"entrypoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ovm_network":    dataSourceOvmNetwork(),
			"ovm_repository": dataSourceOvmRepository(),
			"ovm_serverpool": dataSourceOvmServerPool(),
			"ovm_vm":         dataSourceOvmVM(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"ovm_vm":    resourceOvmVM(),
			"ovm_vmcd":  resourceOvmVmcd(),
			"ovm_vd":    resourceOvmVd(),
			"ovm_vdm":   resourceOvmVdm(),
			"ovm_vmcnm": resourceOvmVmcnm(),
			"ovm_vmcsm": resourceOvmVmcsm(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config Config
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &config); err != nil {
		return nil, err
	}

	log.Println("[INFO] Initializing OVM client")
	return config.Client()
}
