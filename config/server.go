package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Server is the envconfig compatible struct to configure the Web server
type Server struct {
	Scheme      string `envconfig:"WEB_SCHEME" default:"http"`
	BindPort    int    `envconfig:"WEB_BIND_PORT" default:"8080"`
	Host        string `envconfig:"WEB_HOST" default:"localgoprox.com"`
	OutwardPort int    `envconfig:"WEB_PORT" default:"80"`
}

func (s Server) String() string {
	return fmt.Sprintf(
		"scheme: %s, bind_port: %d, host: %s, outward_port: %d",
		s.Scheme,
		s.BindPort,
		s.Host,
		s.OutwardPort,
	)
}

// GetServer gets the Server config using envconfig
func GetServer(appName string) (*Server, error) {
	conf := new(Server)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
