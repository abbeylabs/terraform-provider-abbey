terraform {
  required_providers {
    abbey = {
      source = "abbeylabs/abbey"
    }
  }
}

provider "abbey" {
  bearer_auth = "<token>"
}

resource "abbey_grant_kit" "example" {
  name        = "example_grant_kit"
  description = "Example description."

  workflow = {
    steps = [
      {
        reviewers = {
          one_of = ["replace-me@example.com"]
        }
        skip_if = [
          { bundle = "github://organization/repository/path/to/bundle" }
        ]
      }
    ]
  }

  policies = [
    { bundle = "github://organization/repository/path/to/bundle" }
  ]

  output = {
    location = "github://organization/repo/path/to/access.tf"
    append   = <<-EOT
      resource "abbey_demo" "my_demo" {
        permission = "read_write"
        email = "replace-me@example.com"
      }
    EOT
  }
}
