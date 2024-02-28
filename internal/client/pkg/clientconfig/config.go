package clientconfig

type Config struct {
	BaseUrl     *string
	AccessToken *string
}

func NewConfig() Config {
	baseUrl := "https://api.abbey.io/v1"
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
