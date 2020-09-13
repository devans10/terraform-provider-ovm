package ovm

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceOvmVM() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOvmVMRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repositoryid": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"cpucount": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpucountlimit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"highavailabiltiy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hugepagesenabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memorylimit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ostype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmdomaintype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmmousetype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"osversion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmdiskmappingids": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Computed: true,
				Set:      dataSourceOvmIDHash,
			},
			"virtualnicids": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Computed: true,
				Set:      dataSourceOvmIDHash,
			},
			"serverpoolid": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceOvmVMRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	id, err := client.Vms.GetIDFromName(d.Get("name").(string))
	if err != nil {
		d.SetId("")
		fmt.Println(err)
		return err
	}

	vm, _ := client.Vms.Read(id.Value)

	if vm == nil {
		d.SetId("")
		fmt.Println("Not find any vm")
		return nil
	}

	d.SetId(vm.ID.Value)
	d.Set("name", vm.Name)
	d.Set("cpucount", vm.CPUCount)
	d.Set("cpucountlimit", vm.CPUCountLimit)
	d.Set("highavailability", vm.HighAvailability)
	d.Set("hugepagesenabled", vm.HugePagesEnabled)
	d.Set("memory", vm.Memory)
	d.Set("memorylimit", vm.MemoryLimit)
	d.Set("ostype", vm.OsType)
	d.Set("vmdomaintype", vm.VMDomainType)
	d.Set("vmmousetype", vm.VMMouseType)
	d.Set("osversion", vm.OsVersion)
	d.Set("vmdiskmappingids", flattenIds(vm.VMDiskMappingIds))
	d.Set("virtualnicids", flattenIds(vm.VirtualNicIDs))
	d.Set("repositoryid", vm.RepositoryID)
	d.Set("serverpoolid", vm.ServerPoolID)
	return nil
}

func dataSourceOvmIDHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-",
		strings.ToLower(m["value"].(string))))
	buf.WriteString(fmt.Sprintf("%s-",
		strings.ToLower(m["name"].(string))))
	buf.WriteString(fmt.Sprintf("%s-",
		strings.ToLower(m["type"].(string))))

	return hashcode.String(buf.String())
}
