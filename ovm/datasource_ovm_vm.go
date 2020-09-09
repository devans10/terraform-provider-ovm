package ovm

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/devans10/go-ovm-helper/ovmHelper"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceOvmVm() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOvmVmRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"repositoryid": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"cpucount": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpucountlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"highavailabiltiy": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hugepagesenabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memorylimit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ostype": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmdomaintype": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmmousetype": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"osversion": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmdiskmappingids": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Set:      dataSourceOvmIdHash,
			},
			"virtualnicids": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
				Set:      dataSourceOvmIdHash,
			},
			"serverpoolid": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceOvmVmRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmHelper.Client)

	id, err := client.Vms.GetIdFromName(d.Get("name").(string))
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

	d.SetId(vm.Id.Value)
	d.Set("name", vm.Name)
	d.Set("cpucount", vm.CpuCount)
	d.Set("cpucountlimit", vm.CpuCountLimit)
	d.Set("highavailability", vm.HighAvailability)
	d.Set("hugepagesenabled", vm.HugePagesEnabled)
	d.Set("memory", vm.Memory)
	d.Set("memorylimit", vm.MemoryLimit)
	d.Set("ostype", vm.OsType)
	d.Set("vmdomaintype", vm.VmDomainType)
	d.Set("vmmousetype", vm.VmMouseType)
	d.Set("osversion", vm.OsVersion)
	d.Set("vmdiskmappingids", vm.VmDiskMappingIds)
	d.Set("virtualnicids", vm.VirtualNicIds)
	d.Set("repositoryid", vm.RepositoryId)
	d.Set("serverpoolid", vm.ServerPoolId)
	return nil
}

func dataSourceOvmIdHash(v interface{}) int {
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
