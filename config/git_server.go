package config

import (
	"github.com/kelseyhightower/envconfig"
)

// GitServer is the envconfig compatible struct to configure the Git server
type GitServer struct {
	Scheme   string `envconfig:"GIT_SCHEME" default:"http"`
	BindHost string `envconfig:"GIT_BIND_HOST" default:"localhost"`
	BindPort int    `envconfig:"GIT_BIND_PORT" default:"8081"`
	Host     string `envconfig:"GIT_HOST" default:"localgoprox.com"`
	Port     int    `envconfig:"GIT_PORT" default:"8081"`
}

// GetGitServer gets the GitServer config using envconfig
func GetGitServer(appName string) (*GitServer, error) {
	conf := new(GitServer)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
