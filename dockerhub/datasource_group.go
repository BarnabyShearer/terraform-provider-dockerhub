package dockerhub

import (
	"context"
	"fmt"

	dh "github.com/BarnabyShearer/dockerhub/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The provider specific id of the group.",
				Type:        schema.TypeString,
				Computed:    true,
			},
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
				Computed:    true,
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

func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)

	organisation := d.Get("organisation").(string)
	name := d.Get("name").(string)

	group, err := client.GetGroup(ctx, organisation, name)
	if err != nil {
		return diag.FromErr(err)
	}
	fmt.Println(group)
	d.SetId(fmt.Sprintf("%s/%s", organisation, group.Name))
	d.Set("description", group.Description)
	d.Set("group_id", group.Id)
	return nil
}
