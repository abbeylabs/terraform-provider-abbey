package grantkit_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"abbey.so/terraform-provider-abbey/internal/abbey"
	"abbey.so/terraform-provider-abbey/internal/abbey/provider"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"abbey": providerserver.NewProtocol6WithError(abbey.New("test", provider.DefaultHost)()),
}

func TestAccGrantKit(t *testing.T) {
	randomPostfix := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("acc-test-%s", randomPostfix)

	t.Run("Ok", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ResourceName: "abbey_grant_kit.test",
					Config: fmt.Sprintf(
						`
						resource "abbey_grant_kit" "test" {
						  name = "%s"

						  workflow = {
						    steps = [
						      {
						        reviewers = {
						          one_of = ["primary-id-1"]
						        }
						      }
						    ]
						  }

						#  policies = {
						#	grant_if = [
						#	  { bundle = "github://organization/repository/path/to/bundle.tar.gz" }
						#	]
						#	revoke_if = [
						#	  {
						#		query = <<-EOT
						#		  input.Requester == true
						#		EOT
						#	  }
						#	]
						#  }

						  output = {
						    location = "github://path/to/access.tf"
						    append   = "test"
						  }
						}
						`,
						name,
					),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("abbey_grant_kit.test", "id"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "name", name),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.0", "primary-id-1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.location", "github://path/to/access.tf"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.append", "test"),
						// resource.TestCheckResourceAttrSet("abbey_grant_kit.test", "policies"),
					),
				},
			},
		})
	})
}
