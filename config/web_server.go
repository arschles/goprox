package config

import (
	"github.com/kelseyhightower/envconfig"
)

// WebServer is the envconfig compatible struct to configure the Web server
type WebServer struct {
	Scheme   string `envconfig:"WEB_SCHEME" default:"http"`
	BindHost string `envconfig:"WEB_BIND_HOST" default:"localhost"`
	BindPort int    `envconfig:"WEB_BIND_PORT" default:"8080"`
	Host     string `envconfig:"WEB_HOST" default:"localgoprox.com"`
	Port     int    `envconfig:"WEB_PORT" default:"8080"`
}

// GetWebServer gets the GitServer config using envconfig
func GetWebServer(appName string) (*WebServer, error) {
	conf := new(WebServer)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
