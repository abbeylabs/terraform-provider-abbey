package grantkit_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"abbey.so/terraform-provider-abbey/internal/abbey"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"abbey": providerserver.NewProtocol6WithError(abbey.New("test", "http://localhost:8080")()),
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
						  description = ""

						  workflow = {
							steps = [
							  {
								reviewers = {
								  one_of = ["primary-id-1"]
								}
								skip_if = [
								  { bundle = "github://organization/repository/path/to/bundle.tar.gz" }
								]
							  }
							]
						  }

						  policies = {
							grant_if = [
							  { bundle = "github://organization/repository/path/to/bundle.tar.gz" }
							]
							revoke_if = [
							  {
								query = <<-EOT
								  input.Requester == true
								EOT
							  }
							]
						  }

						  output = {
							location = "github://path/to/access.tf"
							append = <<-EOT
							EOT
						  }
						}
						`,
						name,
					),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("abbey_grant_kit.test", "id"),
						resource.TestCheckResourceAttr("abbey_grant_kit.test", "name", name),
						resource.TestCheckResourceAttrSet("abbey_grant_kit.test", "workflow"),
						resource.TestCheckResourceAttrSet("abbey_grant_kit.test", "policies"),
						resource.TestCheckResourceAttrSet("abbey_grant_kit.test", "output"),
					),
				},
			},
		})
	})
}
