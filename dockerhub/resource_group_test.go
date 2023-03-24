package dockerhub

import (
	"context"
	"fmt"
	"strings"
	"testing"

	dh "github.com/BarnabyShearer/dockerhub/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDockerhubGroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
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
                    `, organisation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("dockerhub_group.foo", "description", "terraform test group"),
				),
			},
		},
	})
}

func testAccCheckGroupResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*dh.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dockerhub_group" {
			continue
		}

		split := strings.Split(rs.Primary.ID, "/")

		if len(split) != 2 {
			return fmt.Errorf("Unexpected Id split: %s", rs.Primary.ID)
		}
		organisation, name := split[0], split[1]

		_, err := client.GetGroup(context.Background(), organisation, name)
		if err == nil {
			return fmt.Errorf("Group (%s) still exists.", rs.Primary.ID)
		}
		if !strings.Contains(err.Error(), "Team not found") {
			return err
		}
	}

	return nil
}
