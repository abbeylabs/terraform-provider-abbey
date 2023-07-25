---
page_title: "Provider: Abbey"
description: |-
    The Abbey provider is used to automate access requests and revocation flows to sensitive data.
---

# Abbey Provider

The `abbey` provider is used to automate access requests and revocation flows to sensitive data.

## Example Usage

```terraform
terraform {
  required_providers {
    abbey = {
      source = "abbeylabs/abbey"
    }
  }
}

provider "abbey" {}

resource "abbey_grant_kit" "example" {
  name        = ""
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
    append   = <<-EOT
    EOT
  }
}
```

## Authentication

TBD
