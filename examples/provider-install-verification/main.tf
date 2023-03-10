terraform {
  required_providers {
    abbey = {
      source = "abbeylabs/abbey"
    }
  }
}

provider "abbey" {}

resource "abbey_grant_kit" "example" {
  name = ""
  description = ""

  workflow = {
    steps = [
      {
        reviewers = {
          one_of = ["primary-id-1"]
        }
        skip_if = [
          { bundle = "github://organization/repository/path/to/bundle.tar.gz" }
        ]
      }
    ]
  }

  policies = {
    grant_if = [
      { bundle = "github://organization/repository/path/to/bundle.tar.gz" }
    ]
    revoke_if = [
      {
        query = <<-EOT
          input.Requester == true
        EOT
      }
    ]
  }

  output = {
    location = "github://organization/repo/path/to/access.tf"
    append = <<-EOT
    EOT
  }
}
