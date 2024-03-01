resource "abbey_grant_kit" "example" {
  name = "name"

  description = "description"

  workflow = {
    steps = [
      {
        reviewers = {
          one_of = [
            "one_of"
          ]
          all_of = [
            "all_of"
          ]
        }

        skip_if = [
          {
            bundle = "bundle"
            query  = "query"
          }

        ]
      }

    ]
  }


  policies = [
    {
      bundle = "bundle"
      query  = "query"
    }

  ]

  output = {
    location  = "location"
    append    = "append"
    overwrite = "overwrite"
  }


}
