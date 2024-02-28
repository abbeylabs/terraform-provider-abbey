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

func getDemoProviderConfig(serverUrl string) string {
	return fmt.Sprintf(`
		provider "abbey" {
        host = "%v"
        auth_token = "auth_token"
}

	`, serverUrl)
}

func setupDemoMockServer() *httptest.Server {
	db := make(map[string]*demo.Demo)
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			Id := path.Base(req.URL.Path)
			demo := db[Id]
			data, _ := json.Marshal(demo)

			rw.Write(data)
			return
		}
		if req.Method == http.MethodPost {
			id := "id"

			demo := demo.Demo{}

			bodyBytes, _ := io.ReadAll(req.Body)
			json.Unmarshal(bodyBytes, &demo)

			demo.SetId(id)
			data, _ := json.Marshal(demo)

			db[id] = &demo

			rw.Write(data)
			return
		}
		if req.Method == http.MethodPatch || req.Method == http.MethodPut {
			demoId := path.Base(req.URL.Path)

			if db[demoId] == nil {
				rw.Write([]byte("Error"))
				return
			}

			demo := demo.Demo{}

			bodyBytes, _ := io.ReadAll(req.Body)
			json.Unmarshal(bodyBytes, &demo)

			demo.SetId(demoId)
			data, _ := json.Marshal(demo)

			db[demoId] = &demo

			rw.Write(data)
			return
		}
	}))

	return server
}

func TestAccabbeyDemoResource(t *testing.T) {
	testAccProtoV6ProviderFactories := map[string]func() (tfprotov6.ProviderServer, error){
		"abbey": providerserver.NewProtocol6WithError(provider.New("test")()),
	}

	server := setupDemoMockServer()
	defer server.Close()
	providerConfig := getDemoProviderConfig(server.URL)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig +
					`
resource "abbey_demo" "example" {
    email = "email"

    permission = "permission"

}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("abbey_demo.example", "email", "email"),
					resource.TestCheckResourceAttr("abbey_demo.example", "permission", "permission"),
				),
			},
		},
	})
}
