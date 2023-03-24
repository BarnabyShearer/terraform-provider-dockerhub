package dockerhub

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDatasourceDockerhubGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "dockerhub_group" "foo" {
						organisation = "%s"
						name = "terraformtest"
						description = "terraform test group"
                    }
                    data "dockerhub_group" "foo" {
						organisation = "%s"
						name = dockerhub_group.foo.name
                    }
                    `, organisation, organisation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.dockerhub_group.foo", "description", "terraform test group"),
				),
			},
		},
	})
}
