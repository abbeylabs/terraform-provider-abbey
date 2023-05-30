package grantkit_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"abbey.io/terraform-provider-abbey/internal/abbey"
	abbeyprovider "abbey.io/terraform-provider-abbey/internal/abbey/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	. "github.com/onsi/gomega"
)

// var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
// 	"abbey": providerserver.NewProtocol6WithError(abbey.New("test", abbeyprovider.DefaultHost)()),
// }

var (
	providerFactories map[string]func() (tfprotov6.ProviderServer, error)
)

func init() {
	providerFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"abbey": providerserver.NewProtocol6WithError(
			abbey.New("test", abbeyprovider.DefaultHost)(),
		),
	}
}

func TestAccGrantKit(t *testing.T) {
	g := NewGomegaWithT(t)
	randomPostfix := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("acc-test-%s", randomPostfix)

	t.Run("Ok", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactories,
			CheckDestroy: func(state *terraform.State) error {
				var res = state.RootModule().Resources["abbey_grant_kit.test"]

				request, err := http.NewRequest(
					http.MethodGet,
					fmt.Sprintf("%s/v1/requestables/%s", abbeyprovider.DefaultHost, res.Primary.ID),
					nil,
				)
				g.Expect(err).NotTo(HaveOccurred())

				request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ABBEY_TOKEN")))

				response, err := http.DefaultClient.Do(request)
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(response.StatusCode).To(Equal(http.StatusNotFound))

				return nil
			},
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(
						`
						resource "abbey_grant_kit" "test" {
						  name        = "%s"
						  description = "test description"
						
						  workflow = {
						    steps = [
						      {
						        reviewers = {
						          one_of = ["primary-id-1"]
						        }
						      }
						    ]
						  }
						
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
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "description", "test description"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.0", "primary-id-1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.location", "github://path/to/access.tf"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.append", "test"),
					),
				},
				// Mutate `name`.
				{
					Config: fmt.Sprintf(
						`
						resource "abbey_grant_kit" "test" {
						  name        = "%s-updated"
						  description = "test description"
						
						  workflow = {
						    steps = [
						      {
						        reviewers = {
						          one_of = ["primary-id-1"]
						        }
						      }
						    ]
						  }
						
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
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "name", fmt.Sprintf("%s-updated", name)),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "description", "test description"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.0", "primary-id-1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.location", "github://path/to/access.tf"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.append", "test"),
					),
				},
				// Mutate a nested attribute.
				{
					Config: fmt.Sprintf(
						`
						resource "abbey_grant_kit" "test" {
						  name        = "%s-updated"
						  description = "test description"
						
						  workflow = {
						    steps = [
						      {
						        reviewers = {
						          one_of = ["different-primary-id"]
						        }
						      }
						    ]
						  }
						
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
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "name", fmt.Sprintf("%s-updated", name)),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "description", "test description"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.0", "different-primary-id"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.location", "github://path/to/access.tf"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.append", "test"),
					),
				},
			},
		})
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(
						`
						resource "abbey_grant_kit" "test" {
						  name        = "%s"
						  description = "test description"

						  workflow = {
						    steps = [
						      {
						        reviewers = {
						          one_of = ["primary-id-1"]
						        }
						      }
						    ]
						  }

						  policies = {
						    grant_if = [
						      { bundle = "github://organization/repository/path/to/bundle.tar.gz" },
						    ]
						    revoke_if = [
						      { query = "input.Requester == true" },
						    ]
						  }

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
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "description", "test description"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "workflow.steps.0.reviewers.one_of.0", "primary-id-1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.location", "github://path/to/access.tf"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "output.append", "test"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "policies.grant_if.#", "1"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "policies.grant_if.0.bundle", "github://organization/repository/path/to/bundle.tar.gz"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "policies.revoke_if.0.query", "input.Requester == true"),
					),
				},
			},
		})
	})
}
