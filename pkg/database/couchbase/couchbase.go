package couchbase

import (
	"context"
	"fintech-api/config"
	"fintech-api/pkg/domain"
	"fintech-api/pkg/repository"
	"github.com/couchbase/gocb/v2"
	"log"
)

type CouchbaseRepo struct {
	client *gocb.Cluster
}

func (cb *CouchbaseRepo) LoginRepository(c context.Context, userName, password string) error {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) RegisterUserRepository(c context.Context, userName, email, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) GetUserRepository(c context.Context, userName string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) GetAccountRepository(c context.Context, userName string) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) AddMoneyRepository(c context.Context, to string, amount int) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) RemoveMoneyRepository(c context.Context, from string, amount int) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) AddToTransaction(c context.Context, from, to string, amount int) error {
	//TODO implement me
	panic("implement me")
}

func NewCouchbaseRepo(cluster *gocb.Cluster, basket string) (repository.FintechRepository, error) {
	return &CouchbaseRepo{
		client: cluster,
	}, nil
}

func InitCouchBase(config config.CouchBaseConfig) (*gocb.Cluster, error) {

	cluster, err := gocb.Connect("couchbase://"+config.Host, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: config.User,
			Password: config.Password,
		},
	})
	if err != nil {
		log.Println("InitCouchBase1: ", err)
		return nil, err
	}

	return cluster, nil
}
