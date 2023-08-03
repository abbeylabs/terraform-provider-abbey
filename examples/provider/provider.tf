terraform {
  required_providers {
    abbey = {
      source  = "abbeylabs/abbey"
      version = "0.2.5"
    }
  }
}

provider "abbey" {
  # Configuration options
  bearer_auth = "<token>"
}
