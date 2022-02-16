package dockerhub

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dh "github.com/BarnabyShearer/dockerhub/v2"
)

func resourceRepositoryGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an organization group / repository permission binding",
		CreateContext: resourceRepositoryGroupCreate,
		UpdateContext: resourceRepositoryGroupUpdate,
		ReadContext:   resourceRepositoryGroupRead,
		DeleteContext: resourceRepositoryGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The provider specific id of the repository group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository path.",
			},
			"group": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The group to add.",
			},
			"groupname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The group to add.",
			},
			"permission": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "The permission to assign the group. One of 'read', 'write' and 'admin'.",
			},
		},
	}
}

func updateRepositoryGroupRequest(d *schema.ResourceData) dh.RepositoryGroup {
	// Validate permission

	return dh.RepositoryGroup{
		GroupId:    d.Get("group").(int),
		GroupId2:   d.Get("group").(int),
		GroupName:  d.Get("groupname").(string),
		GroupName2: d.Get("groupname").(string),
		Permission: d.Get("permission").(string),
	}
}

func resourceRepositoryGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	repository := d.Get("repository").(string)
	repository_group, err := client.CreateRepositoryGroup(ctx, repository, updateRepositoryGroupRequest(d))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%s/%d", repository, repository_group.GroupId))
	return nil
}

func resourceRepositoryGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)

	split := strings.Split(d.Id(), "/")
	if len(split) != 3 {
		return diag.Errorf("Unexpected Id split: %s", d.Id())
	}
	organisation, repo_name, group := split[0], split[1], split[2]
	repository := fmt.Sprintf("%s/%s", organisation, repo_name)

	_, err := client.GetRepositoryGroup(ctx, repository, group)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceRepositoryGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)

	split := strings.Split(d.Id(), "/")
	if len(split) != 3 {
		return diag.Errorf("Unexpected Id split: %s", d.Id())
	}
	organisation, repo_name, group := split[0], split[1], split[2]
	repository := fmt.Sprintf("%s/%s", organisation, repo_name)

	repository_group, err := client.UpdateRepositoryGroup(ctx, repository, group, updateRepositoryGroupRequest(d))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%s/%d", repository, repository_group.GroupId))
	return nil
}

func resourceRepositoryGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)

	split := strings.Split(d.Id(), "/")
	if len(split) != 3 {
		return diag.Errorf("Unexpected Id split: %s", d.Id())
	}
	organisation, repo_name, group := split[0], split[1], split[2]
	repository := fmt.Sprintf("%s/%s", organisation, repo_name)

	err := client.DeleteRepositoryGroup(ctx, repository, group)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
