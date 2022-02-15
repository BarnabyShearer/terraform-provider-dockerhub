terraform {
  required_providers {
    dockerhub = {
      source = "magenta-aps/dockerhub"
    }
  }
}

provider "dockerhub" {
  # Note this cannot be a Personal Access Token
  username = "USERNAME" # optionally use DOCKER_USERNAME env var
  password = "PASSWORD" # optionally use DOCKER_PASSWORD env var
}
