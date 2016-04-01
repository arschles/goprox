package config

import (
	"github.com/kelseyhightower/envconfig"
)

// S3 is the envconfig compatible struct to configure S3 access
type S3 struct {
	Endpoint string `envconfig:"AWS_ENDPOINT" default:"s3.amazonaws.com"`
	Bucket   string `envconfig:"AWS_BUCKET" default:"goprox"`
	Key      string `envconfig:"AWS_KEY"`
	Secret   string `envconfig:"AWS_SECRET"`
}

// GetS3 gets the S3 config using envconfig
func GetS3(appName string) (*S3, error) {
	conf := new(S3)
	if err := envconfig.Process(appName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
