package ovm

import (
	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOvmVmcnm() *schema.Resource {
	return &schema.Resource{
		Create: resourceOvmVmcnmCreate,
		Read:   resourceOvmVmcnmRead,
		Delete: resourceOvmVmcnmDelete,

		//		Update: resourceOvmVmdUpdate,
		/*			Importer: &schema.ResourceImporter{
					State: resourceOvmCheckImporter,
				},*/

		Schema: map[string]*schema.Schema{
			"networkid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vmclonedefinitionid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"virtualnicid": {
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

func checkForResourceVmcnm(d *schema.ResourceData) (ovmhelper.Vmcnm, error) {

	vmcnmParams := &ovmhelper.Vmcnm{}

	// required
	if v, ok := d.GetOk("networkid"); ok {
		vmcnmParams.NetworkID = &ovmhelper.ID{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.Network"}
	}
	if v, ok := d.GetOk("vmclonedefinitionid"); ok {
		vmcnmParams.VMCloneDefinitionID = &ovmhelper.ID{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.VmCloneDefinition"}
	}
	if v, ok := d.GetOk("virtualnicid"); ok {
		vmcnmParams.VirtualNicID = &ovmhelper.ID{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.VirtualNic"}
	}
	if v, ok := d.GetOk("name"); ok {
		vmcnmParams.Name = v.(string)
	}
	return *vmcnmParams, nil
}

func resourceOvmVmcnmRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vmcnm, _ := client.Vmcnms.Read(d.Id())

	if vmcnm == nil {
		d.SetId("")
		return nil
	}

	d.Set("networkid", vmcnm.NetworkID)
	d.Set("vmclonedefinitionid", vmcnm.VMCloneDefinitionID)
	d.Set("virtualnicid", vmcnm.VirtualNicID)
	d.Set("name", vmcnm.Name)
	return nil
}

func resourceOvmVmcnmCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vmcnm, err := checkForResourceVmcnm(d)
	if err != nil {
		return err
	}
	//log.Printf("[INFO] Creating vdm for vmid: %v, vdid: %v, slot: %v", vdm.VmId.Value, vdm.VirtualDiskId.Value, vdm.DiskTarget)

	v, err := client.Vmcnms.Create(vmcnm.VMCloneDefinitionID.Value, vmcnm)
	if err != nil {
		return err
	}

	d.SetId(*v)

	return nil
}

func resourceOvmVmcnmDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)
	//log.Printf("[INFO] Deleting Vdm: %v", d.Id())

	err := client.Vmcnms.Delete(d.Get("vmclonedefinitionid").(string), d.Id())
	if err != nil {
		return err
	}
	return nil
}
