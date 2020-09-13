package ovm

import (
	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOvmVmcsm() *schema.Resource {
	return &schema.Resource{
		Create: resourceOvmVmcsmCreate,
		Read:   resourceOvmVmcsmRead,
		Delete: resourceOvmVmcsmDelete,

		//		Update: resourceOvmVmdUpdate,
		/*			Importer: &schema.ResourceImporter{
					State: resourceOvmCheckImporter,
				},*/

		Schema: map[string]*schema.Schema{
			"vmdiskmappingid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vmclonedefinitionid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repositoryid": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				ForceNew: true,
			},
			"clonetype": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func checkForResourceVmcsm(d *schema.ResourceData) (ovmhelper.Vmcsm, error) {

	vmcsmParams := &ovmhelper.Vmcsm{}

	// required
	if v, ok := d.GetOk("vmdiskmappingid"); ok {
		vmcsmParams.VMDiskMappingID = &ovmhelper.ID{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.VmDiskMapping"}
	}
	if v, ok := d.GetOk("vmclonedefinitionid"); ok {
		vmcsmParams.VMCloneDefinitionID = &ovmhelper.ID{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.VmCloneDefinition"}
	}
	if v, ok := d.GetOk("repositoryid"); ok {
		vmcsmParams.RepositoryID = &ovmhelper.ID{
			Value: v.(map[string]interface{})["value"].(string),
			Type:  v.(map[string]interface{})["type"].(string),
		}
	}
	if v, ok := d.GetOk("clonetype"); ok {
		vmcsmParams.CloneType = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		vmcsmParams.Name = v.(string)
	}
	return *vmcsmParams, nil
}

func resourceOvmVmcsmRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vmcsm, _ := client.Vmcsms.Read(d.Id())

	if vmcsm == nil {
		d.SetId("")
		return nil
	}

	d.Set("vmdiskmappingid", vmcsm.VMDiskMappingID.Value)
	d.Set("vmclonedefinitionid", vmcsm.VMCloneDefinitionID.Value)
	d.Set("repositoryid", flattenID(vmcsm.RepositoryID))
	d.Set("clonetype", vmcsm.CloneType)
	d.Set("name", vmcsm.Name)
	return nil
}

func resourceOvmVmcsmCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vmcsm, err := checkForResourceVmcsm(d)
	if err != nil {
		return err
	}
	//log.Printf("[INFO] Creating vdm for vmid: %v, vdid: %v, slot: %v", vdm.VmId.Value, vdm.VirtualDiskId.Value, vdm.DiskTarget)

	v, err := client.Vmcsms.Create(vmcsm.VMCloneDefinitionID.Value, vmcsm)
	if err != nil {
		return err
	}

	d.SetId(*v)

	return nil
}

func resourceOvmVmcsmDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)
	//log.Printf("[INFO] Deleting Vdm: %v", d.Id())

	err := client.Vmcsms.Delete(d.Get("vmclonedefinitionid").(string), d.Id())
	if err != nil {
		return err
	}
	return nil
}
