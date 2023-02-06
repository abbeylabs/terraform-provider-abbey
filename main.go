package main

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"abbey.so/terraform-provider-abbey/internal/abbey"
)

// GoReleaser sets the version in compiled binaries.
// https://goreleaser.com/cookbooks/using-main.version/
var version = "dev"

const defaultHost = "https://api.abbey.so"

func main() {
	err := providerserver.Serve(
		context.Background(),
		abbey.New(version, defaultHost),
		providerserver.ServeOpts{
			Address:         "registry.terraform.io/abbeylabs/abbey",
			ProtocolVersion: 6,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
