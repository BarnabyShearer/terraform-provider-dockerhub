package dockerhub

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	dh "github.com/BarnabyShearer/dockerhub/v2"
)

func toGenericArray(old []string) []interface{} {
	new := make([]interface{}, len(old))
	for i, v := range old {
		new[i] = v
	}
	return new
}

func TestReadSetString(t *testing.T) {
	cases := []struct {
		input    *schema.Set
		expected []string
	}{
		{
			input: schema.NewSet(
				schema.HashString,
				toGenericArray([]string{"a", "b", "c"}),
			),
			expected: []string{"c", "b", "a"}, // hashed order
		},
		{
			input: schema.NewSet(
				schema.HashString,
				toGenericArray([]string{}),
			),
			expected: []string{},
		},
	}
	for _, c := range cases {
		out := readSetString(c.input)
		if !reflect.DeepEqual(out, c.expected) {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}

func TestAccDockerhubToken_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTokenResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "dockerhub_token" "foo" {
						label = "foo"
						scopes = ["repo:admin"]
					}
					`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("dockerhub_token.foo", "label", "foo"),
				),
			},
			{
				Config: `
					resource "dockerhub_token" "foo" {
						label = "bar"
						scopes = ["repo:admin"]
					}
					`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("dockerhub_token.foo", "label", "bar"),
				),
			},
		},
	})
}

func testAccCheckTokenResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*dh.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dockerhub_token" {
			continue
		}

		_, err := client.GetPersonalAccessToken(context.Background(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Token (%s) still exists.", rs.Primary.ID)
		}
		if !strings.Contains(err.Error(), "does not exist") {
			return err
		}
	}

	return nil
}
