# Terraform Abbey Provider

![GitHub release (with filter)](https://img.shields.io/github/v/release/abbeylabs/terraform-provider-abbey)
![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/abbeylabs/terraform-provider-abbey/release.yaml)
![GitHub issues](https://img.shields.io/github/issues/abbeylabs/terraform-provider-abbey)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B38958%2Fgit%40github.com%3Aabbeylabs%2Fterraform-provider-abbey.git.svg?type=shield)](https://app.fossa.com/projects/custom%2B38958%2Fgit%40github.com%3Aabbeylabs%2Fterraform-provider-abbey.git?ref=badge_shield)
![GitHub](https://img.shields.io/github/license/abbeylabs/terraform-provider-abbey)


The Abbey provider is used by security teams and DevOps/DevSecOps engineers to automate and secure
access to sensitive resources in their existing cloud infrastructure.

Resources in this provider interact with endpoints from the Abbey API and needs to be
configured with credentials before it can be used.

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

{{ tffile "examples/provider-install-verification/main.tf" }}

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

## License

This project is licensed under the MPL 2.0 license. See [LICENSE](LICENSE) for more details.

