package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Git is the envconfig compatible struct to configure the Git server
type Git struct {
	// This is the scheme of the git URL that the go get server tells go get to redirect to
	Scheme string `envconfig:"GIT_SCHEME" default:"http"`
	// This is the host that the go get server will tell go get to redirect to
	Host string `envconfig:"GIT_HOST" default:"git.localgoprox.com"`
}

func (g Git) String() string {
	return fmt.Sprintf("scheme: %s, host: %s", g.Scheme, g.Host)
}

// GetGit gets the Git config using envconfig
func GetGit(appName string) (*Git, error) {
	conf := new(Git)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
