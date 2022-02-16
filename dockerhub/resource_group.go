package dockerhub

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dh "github.com/magenta-aps/dockerhub/v2"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an organization group",
		CreateContext: resourceGroupCreate,
		UpdateContext: resourceGroupUpdate,
		ReadContext:   resourceGroupRead,
		DeleteContext: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"organisation": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The organisation name.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Group name.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "Group description.",
			},
			"group_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Group ID.",
			},
		},
	}
}

func updateGroupRequest(d *schema.ResourceData) dh.Group {
	return dh.Group{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	organisation := d.Get("organisation").(string)
	group, err := client.CreateGroup(ctx, organisation, updateGroupRequest(d))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%s/%s", organisation, group.Name))
	return nil
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)

	split := strings.Split(d.Id(), "/")
	if len(split) != 2 {
		return diag.Errorf("Unexpected Id split: %s", d.Id())
	}
	organisation, name := split[0], split[1]

	group, err := client.GetGroup(ctx, organisation, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("description", group.Description)
	d.Set("group_id", group.GroupId)
	return nil
}

func resourceGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)

	split := strings.Split(d.Id(), "/")
	if len(split) != 2 {
		return diag.Errorf("Unexpected Id split: %s", d.Id())
	}
	organisation, name := split[0], split[1]

	group, err := client.UpdateGroup(ctx, organisation, name, updateGroupRequest(d))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%s/%s", organisation, group.Name))
	return nil
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)

	split := strings.Split(d.Id(), "/")
	if len(split) != 2 {
		return diag.Errorf("Unexpected Id split: %s", d.Id())
	}
	organisation, name := split[0], split[1]

	err := client.DeleteGroup(ctx, organisation, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
