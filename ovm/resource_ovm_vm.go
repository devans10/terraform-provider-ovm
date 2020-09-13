package ovm

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/devans10/go-ovm-helper/ovmhelper"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

type tfVMCfg struct {
	networkID string
}

func resourceOvmVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceOvmVMCreate,
		Read:   resourceOvmVMRead,
		Delete: resourceOvmVMDelete,
		Update: resourceOvmVMUpdate,
		/*			Importer: &schema.ResourceImporter{
					State: resourceOvmCheckImporter,
				},*/

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
				Optional: true,
				Computed: true,
			},
			"repositoryid": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"serverpoolid": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"affinitygroupids": {
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
				Optional: true,
				Set:      dataSourceOvmIDHash,
			},
			"architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bootorder": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"cpucount": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cpucountlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cpupriority": {
				Type:     schema.TypeInt,
				Default:  50,
				Optional: true,
			},
			"cpuutilizationcap": {
				Type:     schema.TypeInt,
				Default:  100,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disklimit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"generation": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"guestdriverversion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"highavailability": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hugepagesenabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"kernelversion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"keymapname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"memorylimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"networkinstallpath": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ostype": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"osverion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pinnedcpus": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"readonly": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"resourcegroupids": {
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
				Optional: true,
				Set:      dataSourceOvmIDHash,
			},
			"restartactiononcrash": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restartactiononpoweroff": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restartactiononrestart": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverid": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
			},
			"sslvncport": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sslttyport": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"userdata": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Optional: true,
				Set:      resourceOvmUserDataHash,
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
				Set:      dataSourceOvmIDHash,
				Optional: true,
				Computed: true,
			},
			"vmapiversion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmclonedefinitionids": {
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
				Set:      dataSourceOvmIDHash,
				Computed: true,
			},
			"vmconfigfileabsolutepath": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmconfigfilemountedpath": {
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
				Set:      dataSourceOvmIDHash,
				Optional: true,
				Computed: true,
			},
			"vmdomaintype": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					switch val.(string) {
					case
						"XEN_HVM",
						"XEN_HVM_PV_DRIVERS",
						"XEN_PVM",
						"UNKNOWN":
						return
					}
					errs = append(errs, fmt.Errorf("%q must be one of [XEN_HVM, XEN_HVM_PV_DRIVERS, XEN_PVM, UNKNOWN], got: %s", key, v))
					return
				},
			},
			"vmmousetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vmrunstate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmstartpolicy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vmclonedefinitionid": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"imageid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"virtualnic": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"networkid": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required: true,
						},
						"macaddress": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vmid": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"generation": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
				Set:      resourceOvmVNHash,
				Optional: true,
				Computed: true,
			},
			"sendmessages": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func checkForResource(d *schema.ResourceData) (ovmhelper.VM, ovmhelper.CfgVM, error) {

	vmParams := &ovmhelper.VM{}
	tfVMCfgParams := &ovmhelper.CfgVM{}

	if v, ok := d.GetOk("name"); ok {
		vmParams.Name = v.(string)
	}
	if v, ok := d.GetOk("repositoryid"); ok {
		vmParams.RepositoryID = &ovmhelper.ID{
			Value: v.(map[string]interface{})["value"].(string),
			Type:  v.(map[string]interface{})["type"].(string),
		}
	}
	if v, ok := d.GetOk("serverpoolid"); ok {
		vmParams.ServerPoolID = &ovmhelper.ID{
			Value: v.(map[string]interface{})["value"].(string),
			Type:  v.(map[string]interface{})["type"].(string),
		}
	}
	if v, ok := d.GetOk("bootorder"); ok {
		vmParams.BootOrder = v.([]string)
	}
	if v, ok := d.GetOk("cpucount"); ok {
		vmParams.CPUCount = v.(int)
	}
	if v, ok := d.GetOk("cpucountlimit"); ok {
		vmParams.CPUCountLimit = v.(int)
	}
	if v, ok := d.GetOk("cpupriority"); ok {
		vmParams.CPUPriority = v.(int)
	}
	if v, ok := d.GetOk("cpuutilizationcap"); ok {
		vmParams.CPUUtilizationCap = v.(int)
	}
	if v, ok := d.GetOk("description"); ok {
		vmParams.Description = v.(string)
	}
	if v, ok := d.GetOk("highavailability"); ok {
		vmParams.HighAvailability = v.(bool)
	}
	if v, ok := d.GetOk("hugepagesenabled"); ok {
		vmParams.HugePagesEnabled = v.(bool)
	}
	if v, ok := d.GetOk("keymapname"); ok {
		vmParams.KeymapName = v.(string)
	}
	if v, ok := d.GetOk("memory"); ok {
		vmParams.Memory = v.(int)
	}
	if v, ok := d.GetOk("memorylimit"); ok {
		vmParams.MemoryLimit = v.(int)
	}
	if v, ok := d.GetOk("networkinstallpath"); ok {
		vmParams.NetworkInstallPath = v.(string)
	}
	if v, ok := d.GetOk("ostype"); ok {
		vmParams.OsType = v.(string)
	}
	if v, ok := d.GetOk("vmdomaintype"); ok {
		vmParams.VMDomainType = v.(string)
	}
	if v, ok := d.GetOk("vmmousetype"); ok {
		vmParams.VMMouseType = v.(string)
	}

	if v, ok := d.GetOk("sendmessages"); ok {
		sendmessages, rootPassword := sendmessagesFromMap(v.(map[string]interface{}))
		tfVMCfgParams.SendMessages = sendmessages
		tfVMCfgParams.RootPassword = rootPassword
	}

	return *vmParams, *tfVMCfgParams, nil
}

func resourceOvmVMCreate(d *schema.ResourceData, meta interface{}) error {
	var v *string

	client := meta.(*ovmhelper.Client)

	vm, tfVMCfgParams, err := checkForResource(d)
	if err != nil {
		return err
	}

	if d.Get("imageid").(string) == "" {
		v, err = client.Vms.CreateVM(vm, tfVMCfgParams)
		if err != nil {
			return err
		}
	} else {
		v, err = client.Vms.CloneVM(d.Get("imageid").(string), d.Get("vmclonedefinitionid").(string), vm, tfVMCfgParams)
		if err != nil {
			return err
		}
	}

	d.SetId(*v)

	return resourceOvmVMRead(d, meta)
}

func resourceOvmVMRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)

	vm, _ := client.Vms.Read(d.Id())

	if vm == nil {
		d.SetId("")
		fmt.Println("Not find any vm")
		return nil
	}

	d.Set("name", vm.Name)
	d.Set("repositoryid", flattenID(vm.RepositoryID))
	d.Set("serverpoolid", flattenID(vm.ServerPoolID))
	d.Set("affinitygroupids", flattenIds(vm.AffinityGroupIDs))
	d.Set("architecture", vm.Architecture)
	d.Set("bootorder", vm.BootOrder)
	d.Set("cpucount", vm.CPUCount)
	d.Set("cpucountlimit", vm.CPUCountLimit)
	d.Set("cpupriority", vm.CPUPriority)
	d.Set("cpuutilizationcap", vm.CPUUtilizationCap)
	d.Set("description", vm.Description)
	d.Set("disklimit", vm.DiskLimit)
	d.Set("generation", vm.Generation)
	d.Set("guestdriverversion", vm.GuestDriverVersion)
	d.Set("highavailability", vm.HighAvailability)
	d.Set("hugepagesenabled", vm.HugePagesEnabled)
	d.Set("kernelversion", vm.KernelVersion)
	d.Set("keymapname", vm.KeymapName)
	d.Set("locked", vm.Locked)
	d.Set("memory", vm.Memory)
	d.Set("memorylimit", vm.MemoryLimit)
	d.Set("networkinstallpath", vm.NetworkInstallPath)
	d.Set("origin", vm.Origin)
	d.Set("ostype", vm.OsType)
	d.Set("osversion", vm.OsVersion)
	d.Set("pinnedcpus", vm.PinnedCpus)
	d.Set("readonly", vm.ReadOnly)
	d.Set("resourcegroupids", flattenIds(vm.ResourceGroupIds))
	d.Set("restartactiononcrash", vm.RestartActionOnCrash)
	d.Set("restartactiononpoweroff", vm.RestartActionOnPowerOff)
	d.Set("restartactiononrestart", vm.RestartActionOnRestart)
	d.Set("vmdomaintype", vm.VMDomainType)
	d.Set("serverid", flattenID(vm.ServerID))
	d.Set("sslvncport", vm.SslVncPort)
	d.Set("sslttyport", vm.SslTtyPort)
	d.Set("userdata", vm.UserData)
	d.Set("virutalnicids", flattenIds(vm.VirtualNicIDs))
	d.Set("vmapiversion", vm.VMApiVersion)
	d.Set("vmclonedefinitions", flattenIds(vm.VMCloneDefinitionIDs))
	d.Set("vmconfigfileabsolutepath", vm.VMConfigFileAbsolutePath)
	d.Set("vmconfigfilemountedpath", vm.VMConfigFileMountedPath)
	d.Set("vmdiskmappingids", flattenIds(vm.VMDiskMappingIds))
	d.Set("vmdomaintype", vm.VMDomainType)
	d.Set("vmmousetype", vm.VMMouseType)
	d.Set("vmrunstate", vm.VMRunState)
	d.Set("vmstartpolicy", vm.VMStartPolicy)

	return nil
}

func resourceOvmVMDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)
	log.Printf("[INFO] Stopping Vm: %v", d.Id())
	err := client.Vms.Stop(d.Id())
	if err != nil {
		return err
	}
	log.Printf("[INFO] Deleting Vm: %v", d.Id())
	err = client.Vms.DeleteVM(d.Id())
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceOvmVMUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmhelper.Client)
	vm, _, err := checkForResource(d)
	if err != nil {
		return err
	}
	err = client.Vms.UpdateVM(d.Id(), vm)
	if err != nil {
		return err
	}
	return resourceOvmVMRead(d, meta)
}

func sendmessagesFromMap(m map[string]interface{}) (*[]ovmhelper.KeyValuePair, *[]ovmhelper.KeyValuePair) {

	result := make([]ovmhelper.KeyValuePair, 0, len(m))
	password := make([]ovmhelper.KeyValuePair, 0, len(m))
	for k, v := range m {
		t := ovmhelper.KeyValuePair{
			Key:   k,
			Value: v.(string),
		}
		if k == "com.oracle.linux.root-password" {
			password = append(password, t)
		} else {
			result = append(result, t)
		}
	}

	return &result, &password
}

func resourceOvmUserDataHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-",
		strings.ToLower(m["value"].(string))))
	buf.WriteString(fmt.Sprintf("%s-",
		strings.ToLower(m["key"].(string))))

	return hashcode.String(buf.String())
}

func resourceOvmVNHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(fmt.Sprintf("%s-",
		strings.ToLower(m["macaddress"].(string))))
	buf.WriteString(fmt.Sprintf("%s-",
		strings.ToLower(m["name"].(string))))

	return hashcode.String(buf.String())
}
