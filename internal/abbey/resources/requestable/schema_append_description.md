String to append to the file.
In addition to normal [Terraform string interpolation][1],
you can also use [Go-style template][1] syntax. 
Go-style template will be rendered when a user submits a resource
request with the following context:

```go
package abbey

type Context struct {
  // Input is the user provided arbitrary input at request time.
  // The context also contains conversion functions for the input,
  // for example, `to_map` converts the input to a map to access
  // nested data.
  Input any

  // Requester contains identity and organizational metadata about
  // the requester.
  Requester struct{
    CanonicalIdentity struct{
      Username string
    }

    // PrimitiveIdentities of the requester. For example, if the user
    // has a GitHub account, you can access it with
    // `.PrimitiveIdentities.Github`.
    PrimitiveIdentities map[string]struct{
      Username string
    }
  }
}
```

A sample template looks like this:

```
append = <<-EOT
  {{"{{"}} $input := .Input | to_map -{{"}}"}}
  resource "github_repository_collaborator" "${trimprefix(each.value.full_name, "abbeylabs/")}_{{"{{"}} $.Requester.CanonicalIdentity.Username {{"}}"}} {
    repository = "${ trimprefix(each.value.full_name, "abbeylabs/") }"
    username   = "{{"{{"}} $.Requester.PrimitiveIdentities.Github.Username {{"}}"}}
    permission = "{{"{{"}} $input.permission {{"}}"}}
  }
EOT
```

[1]: https://developer.hashicorp.com/terraform/language/expressions/strings#string-templates
[2]: https://pkg.go.dev/text/template
