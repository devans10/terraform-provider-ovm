package ovm

import (
	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOvmVmcd() *schema.Resource {
	return &schema.Resource{
		Create: resourceOvmVmcdCreate,
		Read:   resourceOvmVmcdRead,
		Delete: resourceOvmVmcdDelete,

		//		Update: resourceOvmVmdUpdate,
		/*			Importer: &schema.ResourceImporter{
					State: resourceOvmCheckImporter,
				},*/

		Schema: map[string]*schema.Schema{
			"vmid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func checkForResourceVmcd(d *schema.ResourceData) (ovmhelper.Vmcd, error) {

	vmcdParams := &ovmhelper.Vmcd{}

	// required
	if v, ok := d.GetOk("vmid"); ok {
		vmcdParams.VMID = &ovmhelper.ID{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.Vm"}
	}
	if v, ok := d.GetOk("name"); ok {
		vmcdParams.Name = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		vmcdParams.Description = v.(string)
	}
	return *vmcdParams, nil
}

func resourceOvmVmcdRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vmcd, _ := client.Vmcds.Read(d.Id())

	if vmcd == nil {
		d.SetId("")
		return nil
	}

	d.Set("vmid", vmcd.VMID)
	d.Set("description", vmcd.Description)
	d.Set("name", vmcd.Name)
	return nil
}

func resourceOvmVmcdCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vmcd, err := checkForResourceVmcd(d)
	if err != nil {
		return err
	}
	//log.Printf("[INFO] Creating vdm for vmid: %v, vdid: %v, slot: %v", vdm.VmId.Value, vdm.VirtualDiskId.Value, vdm.DiskTarget)

	v, err := client.Vmcds.Create(vmcd.VMID.Value, vmcd)
	if err != nil {
		return err
	}

	d.SetId(*v)

	return nil
}

func resourceOvmVmcdDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)
	//log.Printf("[INFO] Deleting Vdm: %v", d.Id())

	err := client.Vmcds.Delete(d.Get("vmid").(string), d.Id())
	if err != nil {
		return err
	}
	return nil
}
