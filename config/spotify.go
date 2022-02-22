package config

type Spotify struct {
	Enabled bool `yaml:"enabled"`

	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURI  string `yaml:"redirect_uri"`
}
