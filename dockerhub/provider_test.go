package dockerhub

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"dockerhub": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("DOCKER_USERNAME"); v == "" {
		t.Fatal("DOCKER_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("DOCKER_PASSWORD"); v == "" {
		t.Fatal("DOCKER_PASSWORD must be set for acceptance tests")
	}
}
