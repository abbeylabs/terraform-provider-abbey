package identity_test

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

func TestAccIdentity(t *testing.T) {
	randomPostfix := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("acc-test-%s", randomPostfix)

	t.Run("Ok", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
			Steps: []resource.TestStep{
				{
					ResourceName: "abbey_identity.test",
					Config: fmt.Sprintf(
						`
						resource "abbey_identity" "test" {
							name   = "%s"
							linked = jsonencode({
								abbey = [
									{
										type  = "AuthId"
										value = "email@example.com"
									},
								]
								a = [1]
								b = [
									{ prop = true },
									{ prop = false },
								]
							})
						}
						`,
						name,
					),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("abbey_identity.test", "id"),
						resource.TestCheckResourceAttr("abbey_identity.test", "name", name),
						resource.TestCheckResourceAttrSet("abbey_identity.test", "linked"),
					),
				},
			},
		})
	})
}
