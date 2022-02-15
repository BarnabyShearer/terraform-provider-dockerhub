package dockerhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	rtd "github.com/magenta-aps/dockerhub/v2"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOCKER_USERNAME", nil),
				Description: "Username for authentication.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOCKER_PASSWORD", nil),
				Description: "Password for authentication.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dockerhub_repository": resourceRepository(),
			"dockerhub_token":      resourceToken(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return rtd.NewClient(d.Get("username").(string), d.Get("password").(string)), nil
}
