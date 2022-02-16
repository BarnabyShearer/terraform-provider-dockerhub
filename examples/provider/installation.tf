terraform {
  required_version = ">= 0.13"

  required_providers {
    kubectl = {
      source  = "magenta-aps/dockerhub"
      version = ">= 0.0.12"
    }
  }
}
