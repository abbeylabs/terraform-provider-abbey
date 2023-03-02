package main

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name abbey

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"abbey.so/terraform-provider-abbey/internal/abbey"
	"abbey.so/terraform-provider-abbey/internal/abbey/provider"
)

// GoReleaser sets the version in compiled binaries.
// https://goreleaser.com/cookbooks/using-main.version/
var version = "dev"

func main() {
	err := providerserver.Serve(
		context.Background(),
		abbey.New(version, provider.DefaultHost),
		providerserver.ServeOpts{
			Address:         "registry.terraform.io/abbeylabs/abbey",
			ProtocolVersion: 6,
			Debug:           false,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
