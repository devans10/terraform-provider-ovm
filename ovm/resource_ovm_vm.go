package ovm

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/devans10/go-ovm-helper/ovmHelper"
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
				Required: false,
				ForceNew: false,
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
				Required: false,
				ForceNew: false,
			},
			"cpucountlimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"cpupriority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"cpuutilizationcap": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"diskLimit": {
				Type:     schema.TypeInt,
				Optional: true,
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
				Computed: true,
			},
			"hugepagesenabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
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
				Required: false,
				ForceNew: false,
			},
			"memorylimit": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"networkinstallpath": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
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
				Required: false,
				ForceNew: false,
			},
			"restartactiononpoweroff": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"restartactiononrestart": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"serverid": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"sslvncport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"sslttyport": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
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
			"vmapiversion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmclonedefinitionids": {
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
				Set:      dataSourceOvmIDHash,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: true,
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
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"vmmousetype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"vmrunstate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"vmstartpolicy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: false,
			},
			"clonevmid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"networkid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"sendmessages": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"startvm": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func checkForResource(d *schema.ResourceData) (ovmHelper.Vm, ovmHelper.CfgVm, error) {

	vmParams := &ovmHelper.Vm{}
	tfVMCfgParams := &ovmHelper.CfgVm{}

	// required
	if v, ok := d.GetOk("repositoryid"); ok {
		vmParams.RepositoryId = &ovmHelper.Id{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.Repository"}
	}

	if v, ok := d.GetOk("serverpoolid"); ok {
		vmParams.ServerPoolId = &ovmHelper.Id{Value: v.(string),
			Type: "com.oracle.ovm.mgr.ws.model.ServerPool"}
	}

	if v, ok := d.GetOk("vmdomaintype"); ok {
		vmParams.VmDomainType = v.(string)
	}

	//Optinal parameters
	if v, ok := d.GetOk("cpucount"); ok {
		vmParams.CpuCount = v.(int)
	}
	if v, ok := d.GetOk("cpucountlimit"); ok {
		vmParams.CpuCountLimit = v.(int)
	}
	if v, ok := d.GetOk("name"); ok {
		vmParams.Name = v.(string)
	}
	if v, ok := d.GetOk("hugepagesenabled"); ok {
		vmParams.HugePagesEnabled = v.(bool)
	}
	if v, ok := d.GetOk("memory"); ok {
		vmParams.Memory = v.(int)
	}
	if v, ok := d.GetOk("networkid"); ok {
		tfVMCfgParams.NetworkId = v.(string)
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

	client := meta.(*ovmHelper.Client)

	vm, tfVMCfgParams, err := checkForResource(d)
	if err != nil {
		return err
	}

	if d.Get("clonevmid").(string) == "" {
		v, err = client.Vms.CreateVm(vm, tfVMCfgParams)
		if err != nil {
			return err
		}
	} else {
		v, err = client.Vms.CloneVm(d.Get("clonevmid").(string), d.Get("vmclonedefinitionid").(string), vm, tfVMCfgParams)
		if err != nil {
			return err
		}

	}

	if d.Get("startvm").(bool) {
		err = client.Vms.Start(*v)
		if err != nil {
			return err
		}
	}

	d.SetId(*v)

	return nil
}

func resourceOvmVMRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmHelper.Client)

	vm, _ := client.Vms.Read(d.Id())

	if vm == nil {
		d.SetId("")
		fmt.Println("Not find any vm")
		return nil
	}

	d.Set("name", vm.Name)
	d.Set("repositoryid", vm.RepositoryId)
	d.Set("serverpoolid", vm.ServerPoolId)
	d.Set("affinitygroupids", flattenIds(vm.AffinityGroupIds))
	d.Set("architecture", vm.Architecture)
	d.Set("bootorder", vm.BootOrder)
	d.Set("cpucount", vm.CpuCount)
	d.Set("cpucountlimit", vm.CpuCountLimit)
	d.Set("cpupriority", vm.CpuPriority)
	d.Set("cpuutilizationcap", vm.CpuUtilizationCap)
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
	d.Set("vmdomaintype", vm.VmDomainType)
	d.Set("serverid", vm.ServerId)
	d.Set("sslvncport", vm.SslVncPort)
	d.Set("sslttyport", vm.SslTtyPort)
	d.Set("userdata", vm.UserData)
	d.Set("virutalnicids", flattenIds(vm.VirtualNicIds))
	d.Set("vmapiversion", vm.VmApiVersion)
	d.Set("vmclonedefinitions", flattenIds(vm.VmCloneDefinitionIds))
	d.Set("vmconfigfileabsolutepath", vm.VmConfigFileAbsolutePath)
	d.Set("vmconfigfilemountedpath", vm.VmConfigFileMountedPath)
	d.Set("vmdiskmappingids", flattenIds(vm.VmDiskMappingIds))
	d.Set("vmdomaintype", vm.VmDomainType)
	d.Set("vmmousetype", vm.VmMouseType)
	d.Set("vmrunstate", vm.VmRunState)
	d.Set("vmstartpolicy", vm.VmStartPolicy)

	return nil
}

func resourceOvmVMDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmHelper.Client)
	log.Printf("[INFO] Deleting Vm: %v", d.Id())
	err := client.Vms.DeleteVm(d.Id())
	if err != nil {
		return err
	}
	return nil
}

func resourceOvmVMUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ovmHelper.Client)
	vm, _, err := checkForResource(d)
	if err != nil {
		return err
	}
	err = client.Vms.UpdateVm(d.Id(), vm)
	if err != nil {
		return err
	}
	return nil
}

func sendmessagesFromMap(m map[string]interface{}) (*[]ovmHelper.KeyValuePair, *[]ovmHelper.KeyValuePair) {

	result := make([]ovmHelper.KeyValuePair, 0, len(m))
	password := make([]ovmHelper.KeyValuePair, 0, len(m))
	for k, v := range m {
		t := ovmHelper.KeyValuePair{
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
