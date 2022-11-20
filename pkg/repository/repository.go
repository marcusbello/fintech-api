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
	// couchbase bucket
	bucket *gocb.Bucket
}

// AddMoneyRepository :
func (r fintechRepository) AddMoneyRepository(c *gin.Context, to string, amount int) (domain.Account, error) {
	var res domain.Account
	userScope := r.bucket.Scope("bank")
	userCol := userScope.Collection("account")
	userID := strings.ToLower(to)
	result, err := userCol.LookupIn(userID, []gocb.LookupInSpec{
		gocb.GetSpec("balance", nil),
	}, nil)
	if errors.Is(err, gocb.ErrDocumentNotFound) {
		return res, errors.New("user not found")
	} else if err != nil {
		return res, err
	}
	var balance int
	err = result.ContentAt(0, &balance)
	if err != nil {
		return res, err
	}
	// prepare response
	res = domain.Account{
		UserId:  userID,
		Balance: balance + amount,
	}
	_, err = userCol.Upsert(to, &res, nil)
	if err != nil {
		return res, err
	}
	return res, nil
}

// RemoveMoneyRepository :
func (r fintechRepository) RemoveMoneyRepository(c *gin.Context, from string, amount int) (domain.Account, error) {
	var res domain.Account
	userScope := r.bucket.Scope("bank")
	userCol := userScope.Collection("account")
	userID := strings.ToLower(from)
	result, err := userCol.LookupIn(userID, []gocb.LookupInSpec{
		gocb.GetSpec("balance", nil),
	}, nil)
	if errors.Is(err, gocb.ErrDocumentNotFound) {
		return res, errors.New("user not found")
	} else if err != nil {
		return res, err
	}
	var balance int
	err = result.ContentAt(0, &balance)
	if err != nil {
		return res, err
	}
	if balance < amount {
		return domain.Account{}, errors.New("insufficient funds")
	}
	balance = balance - amount
	// prepare response
	res = domain.Account{
		UserId:  userID,
		Balance: balance,
	}
	_, err = userCol.Upsert(userID, &res, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}

// LoginRepository :
func (r fintechRepository) LoginRepository(c *gin.Context, userName, password string) error {
	//var user domain.UserType
	userScope := r.bucket.Scope("bank")
	userCol := userScope.Collection("users")
	result, err := userCol.LookupIn(strings.ToLower(userName), []gocb.LookupInSpec{
		gocb.GetSpec("password", nil),
	}, nil)
	if errors.Is(err, gocb.ErrDocumentNotFound) {
		return errors.New("user not found")
	} else if err != nil {
		return err
	}
	var userPass string
	err = result.ContentAt(0, &userPass)
	if err != nil {
		return err
	}
	if userPass != password {
		return err
	}
	return nil
}

// RegisterUserRepository :
func (r fintechRepository) RegisterUserRepository(c *gin.Context, userName, email, password string) (string, error) {
	var res domain.UserType
	res = domain.UserType{
		UserName: userName,
		Email:    email,
		Password: password,
		Account:  domain.Account{},
	}
	userScope := r.bucket.Scope("users")
	col := userScope.Collection("account")
	_, err := col.Insert(userName, &res, nil)
	if err != nil {
		return "", err
	}

	return res.UserName, nil
}

// GetUserNameRepository :
func (r fintechRepository) GetUserRepository(c *gin.Context, userName string) (domain.UserType, error) {
	//TODO implement me
	panic("implement me")
}

// GetAccountRepository :
func (r fintechRepository) GetAccountRepository(c *gin.Context, userName string) (domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

// TransferMoneyRepository :
func (r fintechRepository) TransferMoneyRepository(c *gin.Context, to, from string, amount int) (domain.Account, error) {
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
