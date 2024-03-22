//go:build unit
// +build unit

package grantkit

import (
	"context"
	"github.com/go-provider-sdk/pkg/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigureResource(t *testing.T) {
	// mock client
	mockClient := &client.Client{}

	// create GrantKitResource instance
	resourceInstance := NewGrantKitResource()

	// Type-assert to the concrete type
	r, ok := resourceInstance.(*GrantKitResource)
	if !ok {
		t.Fatalf("Failed to type assert resourceInstance to *GrantKitResource")
	}

	// create mock ConfigureRequest
	req := resource.ConfigureRequest{
		ProviderData: mockClient,
	}

	var resp resource.ConfigureResponse

	r.Configure(context.Background(), req, &resp)

	// assertions
	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, mockClient, r.client, "Expected client to be set correctly in GrantKitResource")
}
