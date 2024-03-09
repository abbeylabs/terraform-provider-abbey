package client

import (
	"github.com/go-provider-sdk/internal/configmanager"
	"github.com/go-provider-sdk/pkg/clientconfig"
	"github.com/go-provider-sdk/pkg/demo"
	"github.com/go-provider-sdk/pkg/grantkits"
	"github.com/go-provider-sdk/pkg/identities"
)

type Client struct {
	GrantKits  *grantkits.GrantKitsService
	Identities *identities.IdentitiesService
	Demo       *demo.DemoService
	manager    *configmanager.ConfigManager
}

func NewClient(config clientconfig.Config) *Client {
	manager := configmanager.NewConfigManager(config)
	return &Client{
		GrantKits:  grantkits.NewGrantKitsService(manager),
		Identities: identities.NewIdentitiesService(manager),
		Demo:       demo.NewDemoService(manager),
		manager:    manager,
	}
}

func (c *Client) SetBaseUrl(baseUrl string) {
	c.manager.SetBaseUrl(baseUrl)
}

func (c *Client) SetAccessToken(accessToken string) {
	c.manager.SetAccessToken(accessToken)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
