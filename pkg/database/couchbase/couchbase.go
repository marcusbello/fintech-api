package couchbase

import (
	"fintech-api/config"
	"github.com/couchbase/gocb/v2"
)

func InitCouchBase(config config.CouchBaseConfig) (*gocb.Cluster, error) {

	cluster, err := gocb.Connect("couchbase://"+config.Host, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: config.User,
			Password: config.Password,
		},
	})

	return cluster, err
}
