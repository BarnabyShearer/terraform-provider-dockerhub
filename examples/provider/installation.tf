#terraform {
  required_version = ">= 0.13"

  required_providers {
    dockerhub = {
      source  = "BarnabyShearer/dockerhub"
      version = ">= 0.0.12"
    }
  }
}
