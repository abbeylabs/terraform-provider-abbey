package clientconfig

type Config struct {
	BaseUrl     *string
	AccessToken *string
}

func NewConfig() Config {
	baseUrl := DEFAULT_ENVIRONMENT
	newConfig := Config{
		BaseUrl: &baseUrl,
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}

func (c *Config) SetAccessToken(accessToken string) {
	c.AccessToken = &accessToken
}
