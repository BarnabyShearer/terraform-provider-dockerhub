package dockerhub

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	dh "github.com/magenta-aps/dockerhub/v2"
)

func TestAccDockerhubRepository_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRepositoryResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "dockerhub_repository" "foo" {
						namespace = "barnabyshearer"
						name = "foo"
						description = "bar"
					}
					`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("dockerhub_repository.foo", "description", "bar"),
				),
			},
			{
				Config: `
					resource "dockerhub_repository" "foo" {
						namespace = "barnabyshearer"
						name = "foo"
						description = "baz"
					}
					`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("dockerhub_repository.foo", "description", "baz"),
				),
			},
		},
	})
}

func testAccCheckRepositoryResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*dh.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dockerhub_repository" {
			continue
		}

		_, err := client.GetRepository(context.Background(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Repository (%s) still exists.", rs.Primary.ID)
		}
		if !strings.Contains(err.Error(), "Object not found") {
			return err
		}
	}

	return nil
}
