resource "dockerhub_repository" "example" {
  name             = "example"
  namespace        = "organization_name"
  description      = "Example repository."
  full_description = "Readme."
}
