package ovm

import (
	"log"

	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOvmVdm() *schema.Resource {
	return &schema.Resource{
		Create: resourceOvmVdmCreate,
		Read:   resourceOvmVdmRead,
		Delete: resourceOvmVdmDelete,

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
			"vdid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"slot": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
		},
	}
}

func checkForResourceVdm(d *schema.ResourceData) (ovmhelper.Vdm, error) {

	vdmParams := &ovmhelper.Vdm{}

	// required
	if v, ok := d.GetOk("vmid"); ok {
		vdmParams.VMID = &ovmhelper.ID{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.Vm"}
	}
	if v, ok := d.GetOk("vdid"); ok {
		vdmParams.VirtualDiskID = &ovmhelper.ID{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.VirtualDisk"}
	}
	if v, ok := d.GetOk("slot"); ok {
		vdmParams.DiskTarget = v.(int)
		log.Printf("[DEBUG] Slot: %v DiskTarget: %v", v.(int), vdmParams.DiskTarget)
	}
	//optional
	if v, ok := d.GetOk("description"); ok {
		vdmParams.Description = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		vdmParams.Name = v.(string)
	}

	return *vdmParams, nil
}

func resourceOvmVdmRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vdm, _ := client.Vdms.Read(d.Get("vmid").(string), d.Id())

	if vdm == nil {
		d.SetId("")
		return nil
	}

	d.Set("vmid", vdm.VMID)
	d.Set("vdid", vdm.VirtualDiskID)
	d.Set("slot", vdm.DiskTarget)
	d.Set("description", vdm.Description)
	d.Set("name", vdm.Name)
	return nil
}

func resourceOvmVdmCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vdm, err := checkForResourceVdm(d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Creating vdm for vmid: %v, vdid: %v, slot: %v", vdm.VMID.Value, vdm.VirtualDiskID.Value, vdm.DiskTarget)

	v, err := client.Vdms.Create(vdm)
	if err != nil {
		return err
	}

	d.SetId(*v)

	return nil
}

func resourceOvmVdmDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)
	//log.Printf("[INFO] Deleting Vdm: %v", d.Id())

	err := client.Vdms.Delete(d.Get("vmid").(string), d.Id())
	if err != nil {
		return err
	}
	return nil
}
