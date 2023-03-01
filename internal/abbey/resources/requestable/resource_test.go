package requestable_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"abbey.so/terraform-provider-abbey/internal/abbey"
	"abbey.so/terraform-provider-abbey/internal/abbey/provider"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"abbey": providerserver.NewProtocol6WithError(abbey.New("test", provider.DefaultHost)()),
}

func testAccPreCheck(*testing.T) {
}

func TestAccRequestable(t *testing.T) {
	randomPostfix := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("acc-test-%s", randomPostfix)

	t.Run("Ok", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:                 func() { testAccPreCheck(t) },
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ResourceName: "abbey_requestable.test",
					Config: fmt.Sprintf(
						`
						resource "abbey_requestable" "test" {
							name     = "%s"
							workflow = {
								builtin = {
									one_of = {
										reviewers = []
									}
								}
							}
							grant = {
								generate = {
									github = {
										repo   = "owner/repo"
										path   = "file.tf"
										append = tolist(toset(["a", "a", "b"])).0
									}
								}
							}
						}
						`,
						name,
					),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("abbey_requestable.test", "id"),
						resource.TestCheckResourceAttr("abbey_requestable.test", "name", name),
						resource.TestCheckResourceAttr("abbey_requestable.test", "grant.%", "1"),
					),
				},
			},
		})
	})
}
