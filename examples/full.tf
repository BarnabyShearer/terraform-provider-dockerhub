# Configure the Docker Hub Provider
provider "dockerhub" {
  username = "azurediamond"
  password = "hunter2"
}

# Create an organization group for developers
resource "dockerhub_group" "project-developers" {
  organisation = "organisation"
  name         = "project"
  description  = "Project developers"
}

# Create an organization group for CI
resource "dockerhub_group" "project-ci" {
  organisation = "organisation"
  name         = "ci"
  description  = "Project CI"
}

# Create an image registry
resource "dockerhub_repository" "project" {
  name             = "project"
  namespace        = "organisation"
  description      = "Project description"
}

# Associate our developers group with the registry
resource "dockerhub_repositorygroup" "project-developers" {
  repository = dockerhub_repository.project.id
  group = dockerhub_group.project-developers.group_id
  groupname = dockerhub_group.project-developers.name
  permission = "admin"
}

# Associate our CI group with the registry
resource "dockerhub_repositorygroup" "project-ci" {
  repository = dockerhub_repository.project.id
  group = dockerhub_group.project-ci.group_id
  groupname = dockerhub_group.project-ci.name
  permission = "write"
}
