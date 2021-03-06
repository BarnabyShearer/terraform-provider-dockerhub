---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dockerhub_repository Resource - terraform-provider-dockerhub"
subcategory: ""
description: |-
  A hub.docker.com repository.
---

# dockerhub_repository (Resource)

A hub.docker.com repository.

## Example Usage

```terraform
resource "dockerhub_repository" "example" {
  name             = "example"
  description      = "Example repository."
  full_description = "Readme."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Repository name.
- **namespace** (String) Repository namespace.

### Optional

- **description** (String) Repository description.
- **full_description** (String) Repository full description.
- **private** (Boolean) Is the repository private.

### Read-Only

- **id** (String) The namespace/name of the repository.

## Import

Import is supported using the following syntax:

```shell
# import using the namespace/name
terraform import dockerhub_repository example username/example
```
