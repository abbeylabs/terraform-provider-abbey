resource "abbey_grant_kit" "my_grant_kit" {
  name        = "my_grant_kit"
  description = "My description."

  workflow = {
    steps = [
      {
        reviewers = {
          one_of = ["alice@example.com"]
        }
      },
      {
        reviewers = {
          all_of = ["bob@example.com", "carol@example.com"]
        }
        skip_if = [
          { bundle = "github://my-org/my-repo/policies/on-call-overrides" }
        ]
      }
    ]
  }

  policies = [
    { bundle = "github://my-org/my-repo/policies/common" }
  ]

  output = {
    location = "github://my-org/my-repo/access.tf"
    append   = <<-EOT
      resource "abbey_demo" "my_demo" {
        permission = "read_write"
        email = "alice@example.com"
      }
    EOT
  }
}