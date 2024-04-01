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
	"abbey/v2/internal/shared/models/grant"
	"abbey/v2/internal/shared/models/grant_workflow"
	"abbey/v2/internal/shared/models/output"
	"abbey/v2/internal/shared/models/policy"
	"abbey/v2/internal/shared/models/request"
	"abbey/v2/internal/shared/models/review"
	"abbey/v2/internal/shared/models/reviewers"
	"abbey/v2/internal/shared/models/step"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func getGrantKitProviderConfig(serverUrl string) string {
	return fmt.Sprintf(`
		provider "abbey" {
        server_url = "server_url"
        bearer_auth = "bearer_auth"
}

	`, serverUrl)
}

func setupGrantKitMockServer() *httptest.Server {
	db := make(map[string]*grant_kit.GrantKit)
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			Id := path.Base(req.URL.Path)
			grant_kit := db[Id]
			data, _ := json.Marshal(grant_kit)

			rw.Write(data)
			return
		}
		if req.Method == http.MethodPost {
			id := "id"

			grant_kit := grant_kit.GrantKit{}

			bodyBytes, _ := io.ReadAll(req.Body)
			json.Unmarshal(bodyBytes, &grant_kit)

			grant_kit.SetId(id)
			data, _ := json.Marshal(grant_kit)

			db[id] = &grant_kit

			rw.Write(data)
			return
		}
		if req.Method == http.MethodPatch || req.Method == http.MethodPut {
			grant_kitId := path.Base(req.URL.Path)

			if db[grant_kitId] == nil {
				rw.Write([]byte("Error"))
				return
			}

			grant_kit := grant_kit.GrantKit{}

			bodyBytes, _ := io.ReadAll(req.Body)
			json.Unmarshal(bodyBytes, &grant_kit)

			grant_kit.SetId(grant_kitId)
			data, _ := json.Marshal(grant_kit)

			db[grant_kitId] = &grant_kit

			rw.Write(data)
			return
		}
	}))

	return server
}

func TestAccabbeyGrantKitResource(t *testing.T) {
	testAccProtoV6ProviderFactories := map[string]func() (tfprotov6.ProviderServer, error){
		"abbey": providerserver.NewProtocol6WithError(provider.New("test")()),
	}

	server := setupGrantKitMockServer()
	defer server.Close()
	providerConfig := getGrantKitProviderConfig(server.URL)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: providerConfig +
					`
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
                query = "query"
}

            ]
}

            ]
}


    policies = [
        {
                bundle = "bundle"
                query = "query"
}

    ]

    output = {
                location = "location"
                append = "append"
                overwrite = "overwrite"
}


}

`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Extend this based on the model attributes
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "grant_kit_id_or_name", "grant_kit_id_or_name"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "name", "name"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "description", "description"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "workflow.steps.0.reviewers.one_of.0", "one_of"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "workflow.steps.0.reviewers.all_of.0", "all_of"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "workflow.steps.0.skip_if.0.bundle", "bundle"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "workflow.steps.0.skip_if.0.query", "query"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "policies.0.bundle", "bundle"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "policies.0.query", "query"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "output.location", "location"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "output.append", "append"),
					resource.TestCheckResourceAttr("abbey_grant_kit.example", "output.overwrite", "overwrite"),
				),
			},
		},
	})
}
