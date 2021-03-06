---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dockerhub_token Resource - terraform-provider-dockerhub"
subcategory: ""
description: |-
  A hub.docker.com personal access token (for uploading images).
---

# dockerhub_token (Resource)

A hub.docker.com personal access token (for uploading images).

## Example Usage

```terraform
resource "dockerhub_token" "example" {
  label = "example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **label** (String) Token label.
- **scopes** (Set of String) Permissions e.g. 'repo:admin'

### Read-Only

- **id** (String) The UUID of the token.
- **token** (String, Sensitive) Token to use as password


