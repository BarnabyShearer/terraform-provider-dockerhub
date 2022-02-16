resource "dockerhub_repositorygroup" "example" {
  repository = "organisation/project"
  group      = 123
  groupname  = "groupname"
  permission = "read"
}
