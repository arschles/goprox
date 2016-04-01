package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	AWSEndpoint string `envconfig:"AWS_ENDPOINT" default:"s3.amazonaws.com"`
	AWSBucket   string `envconfing:"AWS_BUCKET" default:"goprox"`
	AWSKey      string `envconfig:"AWS_KEY"`
	AWSSecret   string `envconfig:"AWS_SECRET"`
}

func getConfig(appName string) (*config, error) {
	conf := new(config)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
