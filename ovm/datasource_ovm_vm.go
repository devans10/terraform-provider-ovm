package ovm

import (
	"fmt"

	"github.com/devans10/go-ovm-helper/ovmHelper"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceOvmVm() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceOvmVmRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cpucount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpucountlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"highavailabiltiy": &schema.Schema{
				Type:	  schema.TypeBool,
				Optional: true,
			},
			"hugepagesenabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"memorylimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ostype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmdomaintype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmmousetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"osversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmdiskmappingids": &schema.Schema{
				Type:	  schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
								 Type:	   schema.TypeString,
								 Required: true,
						},
						"value": {
								  Type:		schema.TypeString,
								  Required: true,
						},
						"name": {
								 Type:		schema.TypeString,
								 Required: true,
						},
						"uri": {
								Type:		schema.TypeString,
								Required: true,
						},
					},
				},
			},
			"virtualnicids": &schema.Schema{
				Type:	  schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
								 Type:	   schema.TypeString,
								 Required: true,
						},
						"value": {
								  Type:		schema.TypeString,
								  Required: true,
						},
						"name": {
								 Type:		schema.TypeString,
								 Required: true,
						},
						"uri": {
								Type:		schema.TypeString,
								Required: true,
						},
					},
				},
			},
			"repositoryid": &schema.Schema{
				Type:     schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
			},
			"serverpoolid": &schema.Schema{
				Type:     schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: false,
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