package repository

import (
	"errors"
	"fintech-api/pkg/domain"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type fintechRepository struct {
	//couchbase bucket
	bucket *gocb.Bucket
}

func (f fintechRepository) LoginRepository(c *gin.Context, userName, password string) error {
	//var user domain.UserType
	userScope := f.bucket.Scope("users")
	userCol := userScope.Collection("account")
	result, err := userCol.LookupIn(strings.ToLower(userName), []gocb.LookupInSpec{
		gocb.GetSpec("password", nil),
	}, nil)
	if errors.Is(err, gocb.ErrDocumentNotFound) {
		return errors.New("user not found")
	} else if err != nil {
		return err
	}
	err = result.ContentAt(0, &password)
	if err != nil {
		return err
	}
	var userPass string
	if userPass != password {
		return err
	}
	return nil
}

func (f fintechRepository) RegisterUserRepository(c *gin.Context, userName, email, password string) (string, error) {
	var res domain.UserType
	res = domain.UserType{
		UserName: userName,
		Email:    email,
		Password: password,
		Account:  domain.Account{},
	}
	userScope := f.bucket.Scope("users")
	col := userScope.Collection("account")
	_, err := col.Insert(userName, &res, nil)
	if err != nil {
		return "", err
	}

	return res.UserName, nil
}

func (f fintechRepository) GetUserNameRepository(c *gin.Context, userName string) (domain.UserType, error) {
	//TODO implement me
	panic("implement me")
}

func (f fintechRepository) GetAccountRepository(c *gin.Context, userName string) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (f fintechRepository) TransferMoneyRepository(c *gin.Context, to, from string, amount int) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

// NewFintechRepository :
func NewFintechRepository(cluster *gocb.Cluster, bucketName string) (domain.FintechRepository, error) {
	bucket := cluster.Bucket(bucketName)
	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, err
	}
	return &fintechRepository{bucket: bucket}, nil
}
