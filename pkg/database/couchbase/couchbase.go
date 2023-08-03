package couchbase

import (
	"fintech-api/config"
	"fintech-api/pkg/domain"
	"fintech-api/pkg/repository"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"log"
)

type CouchbaseRepo struct {
	client *gocb.Cluster
}

func (cb *CouchbaseRepo) LoginRepository(c *gin.Context, userName, password string) error {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) RegisterUserRepository(c *gin.Context, userName, email, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) GetUserRepository(c *gin.Context, userName string) (domain.UserType, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) GetAccountRepository(c *gin.Context, userName string) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) AddMoneyRepository(c *gin.Context, to string, amount int) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) RemoveMoneyRepository(c *gin.Context, from string, amount int) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (cb *CouchbaseRepo) AddToTransaction(c *gin.Context, from, to string, amount int) error {
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
