package dockerhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dh "github.com/BarnabyShearer/dockerhub/v2"
)

func resourceToken() *schema.Resource {
	return &schema.Resource{
		Description:   "A hub.docker.com personal access token (for uploading images).",
		CreateContext: resourceTokenCreate,
		ReadContext:   noop,
		DeleteContext: resourceTokenDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The UUID of the token.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"label": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Token label.",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Token to use as password",
			},
			"scopes": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "Permissions e.g. 'repo:admin'",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSetString(set *schema.Set) []string {
	ret := make([]string, len(set.List()))
	for i, raw := range set.List() {
		ret[i] = raw.(string)
	}
	return ret
}

func resourceTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	token, err := client.CreatePersonalAccessToken(ctx, dh.CreatePersonalAccessToken{
		TokenLabel: d.Get("label").(string),
		Scopes:     readSetString(d.Get("scopes").(*schema.Set)),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(token.UUID)
	d.Set("token", token.Token)
	return nil
}

func noop(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	err := client.DeletePersonalAccessToken(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
