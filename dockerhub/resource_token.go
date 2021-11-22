package dockerhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dh "github.com/BarnabyShearer/dockerhub/v2"
)

func resourceToken() *schema.Resource {
	return &schema.Resource{
		Description:   "A hub.docker.io personal access token (for uploading images).",
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
				Elem:        schema.TypeString,
			},
		},
	}
}

func resourceTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	scopesRaw := d.Get("scopes").(*schema.Set).List()
	scopes := make([]string, len(scopesRaw))
	for i, raw := range scopesRaw {
		scopes[i] = raw.(string)
	}
	token, err := client.CreatePersonalAccessToken(ctx, dh.CreatePersonalAccessToken{
		TokenLabel: d.Get("label").(string),
		Scopes:     scopes,
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
