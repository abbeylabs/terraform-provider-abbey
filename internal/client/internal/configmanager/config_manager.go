package configmanager

import "github.com/go-provider-sdk/pkg/clientconfig"

type ConfigManager struct {
	GrantKits  clientconfig.Config
	Identities clientconfig.Config
	Demo       clientconfig.Config
}

func NewConfigManager(config clientconfig.Config) *ConfigManager {
	return &ConfigManager{
		GrantKits:  config,
		Identities: config,
		Demo:       config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.GrantKits.SetBaseUrl(baseUrl)
	c.Identities.SetBaseUrl(baseUrl)
	c.Demo.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.GrantKits.SetAccessToken(accessToken)
	c.Identities.SetAccessToken(accessToken)
	c.Demo.SetAccessToken(accessToken)
}

func (c *ConfigManager) UpdateAccessToken(originalValue string, newValue string) {

	if c.GrantKits.AccessToken != nil && *c.GrantKits.AccessToken == originalValue {
		c.GrantKits.SetAccessToken(newValue)
	}

	if c.Identities.AccessToken != nil && *c.Identities.AccessToken == originalValue {
		c.Identities.SetAccessToken(newValue)
	}

	if c.Demo.AccessToken != nil && *c.Demo.AccessToken == originalValue {
		c.Demo.SetAccessToken(newValue)
	}
}

func (c *ConfigManager) GetGrantKits() *clientconfig.Config {
	return &c.GrantKits
}
func (c *ConfigManager) GetIdentities() *clientconfig.Config {
	return &c.Identities
}
func (c *ConfigManager) GetDemo() *clientconfig.Config {
	return &c.Demo
}
