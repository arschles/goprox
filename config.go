package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	BindHost string `envconfing:"BIND_HOST" default:"localhost"`
	WebPort  int    `envconfig:"WEB_PORT" default:"8080"`
	GitPort  int    `envconfig:"GIT_PORT" default:"8081"`
}

func getConfig(appName string) (*config, error) {
	conf := new(config)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
