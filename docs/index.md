---
page_title: "Provider: Abbey"
description: |-
    The Abbey provider is used by security teams and DevOps/DevSecOps engineers to automate and secure
    access to sensitive resources.
---

# Abbey Provider

The Abbey provider is used by security teams and DevOps/DevSecOps engineers to automate and secure
access to sensitive resources in their existing cloud infrastructure.

Resources in this provider interact with endpoints from the Abbey API and needs to be
configured with credentials before it can used.

## Background

Abbey is a platform for security teams and DevOps/DevSecOps engineers to
automate and secure access to sensitive resources.

With Abbey, you can improve your security and compliance programs by automatically controlling
and right-sizing permissions so the risks around unauthorized access is limited in the event of a breach.

You do this by leveraging your normal Infrastructure as Code (IaC) tooling —
Terraform, CI/CD systems, secrets managers, and Terraform wrappers.
Abbey provides a Terraform plugin and platform that extends your Infrastructure as Code capabilities
from Infrastructure to Identity, Access, and Management (IAM).

## Contributing

Development happens in the [GitHub repo](https://github.com/abbeylabs/terraform-provider-abbey):

- [Releases](https://github.com/abbeylabs/terraform-provider-abbey/releases)
- [Issues](https://github.com/abbeylabs/terraform-provider-abbey/issues)

## Example Usage

```terraform
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
```

## Authentication

Abbey provides the following methods for authenticating to the Abbey API:

- Abbey Token

### Generating Tokens

You can generate an Abbey Token by:

- [Abbey Developers Dashboard](https://app.abbey.io/developers) -> New API Key -> Create

### Using Tokens

Configure your Abbey `provider` block with a `bearer_auth` string.

```hcl
provider "abbey" {
  bearer_auth = "<token>"
}
```

For better security practice, you should place your token into a variable in your `variables.tf` file and set it
using an environment variable or secret store.
