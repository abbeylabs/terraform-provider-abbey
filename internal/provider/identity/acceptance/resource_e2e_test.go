//go:build acceptance
// +build acceptance

package acceptance

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"abbey/v2/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func getIdentityProviderConfig(serverUrl string) string {
	return fmt.Sprintf(`
		provider "abbey" {
        host = "%v"
        auth_token = "auth_token"
}

	`, serverUrl)
}

func setupIdentityMockServer() *httptest.Server {
	db := make(map[string]*identity.Identity)
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			Id := path.Base(req.URL.Path)
			identity := db[Id]
			data, _ := json.Marshal(identity)

			rw.Write(data)
			return
		}
		if req.Method == http.MethodPost {
			id := "id"

			identity := identity.Identity{}

			bodyBytes, _ := io.ReadAll(req.Body)
			json.Unmarshal(bodyBytes, &identity)

			identity.SetId(id)
			data, _ := json.Marshal(identity)

			db[id] = &identity

			rw.Write(data)
			return
		}
		if req.Method == http.MethodPatch || req.Method == http.MethodPut {
			identityId := path.Base(req.URL.Path)

			if db[identityId] == nil {
				rw.Write([]byte("Error"))
				return
			}

			identity := identity.Identity{}

			bodyBytes, _ := io.ReadAll(req.Body)
			json.Unmarshal(bodyBytes, &identity)

			identity.SetId(identityId)
			data, _ := json.Marshal(identity)

			db[identityId] = &identity

			rw.Write(data)
			return
		}
	}))

	return server
}

func TestAccabbeyIdentityResource(t *testing.T) {
	testAccProtoV6ProviderFactories := map[string]func() (tfprotov6.ProviderServer, error){
		"abbey": providerserver.NewProtocol6WithError(provider.New("test")()),
	}

	server := setupIdentityMockServer()
	defer server.Close()
	providerConfig := getIdentityProviderConfig(server.URL)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig +
					`
resource "abbey_identity" "example" {
    abbey_account = "abbey_account"

    source = "source"

    metadata = "metadata"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("abbey_identity.example", "identity_id", "identity_id"),
					resource.TestCheckResourceAttr("abbey_identity.example", "abbey_account", "abbey_account"),
					resource.TestCheckResourceAttr("abbey_identity.example", "source", "source"),
					resource.TestCheckResourceAttr("abbey_identity.example", "metadata", "metadata"),
				),
			},
		},
	})
}
