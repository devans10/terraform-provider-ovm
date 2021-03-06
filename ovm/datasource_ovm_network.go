package ovm

import (
	"fmt"

	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceOvmNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOvmNetworkRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOvmNetworkRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	id, err := client.Network.GetIDFromName(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		fmt.Println(err)
		return err
	}

	d.SetId(id.Value)
	d.Set("value", id.Value)
	d.Set("name", id.Name)
	d.Set("uri", id.URI)
	d.Set("type", id.Type)

	return nil
}
