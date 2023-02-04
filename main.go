package main

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
		providerserver.ServeOpts{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
