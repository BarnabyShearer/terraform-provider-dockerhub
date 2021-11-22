package dockerhub

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dh "github.com/BarnabyShearer/dockerhub/v2"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		Description:   "A hub.docker.io repository.",
		CreateContext: resourceRepositoryCreate,
		UpdateContext: resourceRepositoryUpdate,
		ReadContext:   resourceRepositoryRead,
		DeleteContext: resourceRepositoryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The namespace/name of the repository.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Repository namespace.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Repository name.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Description: "Repository name.",
			},
			"full_description": {
				Type:        schema.TypeString,
				Required:    false,
				Description: "Repository name.",
			},
			"private": {
				Type:        schema.TypeBool,
				Required:    false,
				Default:     false,
				Description: "Is the repository private.",
			},
		},
	}
}

func updateReqest(d *schema.ResourceData) dh.Repository {
	return dh.Repository{
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		FullDescription: d.Get("full_description").(string),
		Private:         d.Get("private").(bool),
	}
}

func resourceRepositoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	repository, err := client.CreateRepository(ctx, updateReqest(d))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%s/%s", repository.Namespace, repository.Name))
	return nil
}

func resourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	repository, err := client.GetRepository(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("description", repository.Description)
	d.Set("full_description", repository.FullDescription)
	d.Set("private", repository.Private)
	return nil
}

func resourceRepositoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	err := client.UpdateRepository(ctx, d.Id(), updateReqest(d))
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceRepositoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	err := client.DeleteRepository(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
