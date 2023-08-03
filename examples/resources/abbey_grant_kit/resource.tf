resource "abbey_grant_kit" "my_grantkit" {
    description = "...my_description..."
            name = "Dallas Kassulke"
            output = {
        append = "...my_append..."
        location = "...my_location..."
        overwrite = "...my_overwrite..."
    }
            policies = [
        {
            bundle = "...my_bundle..."
            query = "...my_query..."
        },
    ]
            workflow = {
        steps = [
            {
                reviewers = {
                    all_of = [
                        "...",
                    ]
                    one_of = [
                        "...",
                    ]
                }
                skip_if = [
                    {
                        bundle = "...my_bundle..."
                        query = "...my_query..."
                    },
                ]
            },
        ]
    }
        }