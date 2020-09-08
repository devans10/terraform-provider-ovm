package ovm

import (
	"fmt"

	"github.com/devans10/go-ovm-helper/ovmHelper"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceOvmRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOvmRepositoryRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOvmRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmHelper.Client)

	id, err := client.Repos.GetIdFromName(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		fmt.Println(err)
		return err
	}

	d.SetId(id.Value)
	d.Set("name", id.Name)
	d.Set("uri", id.Uri)
	d.Set("type", id.Type)

	return nil
}