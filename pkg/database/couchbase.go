package database

import (
	"fintech-api/utils"
	"github.com/couchbase/gocb/v2"
)

func InitDatabase(config utils.CouchDbConfig) (*gocb.Cluster, error) {

	cluster, err := gocb.Connect("couchbase://"+config.Host, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: config.User,
			Password: config.Password,
		},
	})

	return cluster, err
}
