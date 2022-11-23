package repository

import (
	"errors"
	"fintech-api/pkg/domain"
	"github.com/couchbase/gocb/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type fintechRepository struct {
	// couchbase bucket
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

// TransferMoneyRepository :
func (r fintechRepository) TransferMoneyRepository(c *gin.Context, from, to string, amount int) (domain.Account, error) {
	//remove Money
	log.Println("Stage 1", from, to, amount)
	senderAcct, err := r.RemoveMoneyRepository(c, from, amount)
	if err != nil {
		log.Println("RemoveMoney", err)
		return domain.Account{}, err
	}
	log.Println("add money:", senderAcct)
	//add money
	receiverAcct, err := r.AddMoneyRepository(c, to, amount)
	if err != nil {
		log.Println("AddMoney", err)
		return domain.Account{}, err
	}
	log.Println("remove money:", receiverAcct)
	//add to txDocument
	if err = r.AddToTransaction(c, from, to, amount); err != nil {
		log.Println("AddToTransaction", err)
		return domain.Account{}, err
	}
	// attach to response
	return senderAcct, nil
}

// AddMoneyRepository :
func (r fintechRepository) AddMoneyRepository(c *gin.Context, to string, amount int) (domain.Account, error) {
	var res domain.Account
	userScope := r.bucket.Scope("bank")
	userCol := userScope.Collection("accounts")
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
		UserName: userID,
		Balance:  balance + amount,
	}
	_, err = userCol.Upsert(userID, &res, nil)
	if err != nil {
		return res, err
	}
	return res, nil
}

// RemoveMoneyRepository :
func (r fintechRepository) RemoveMoneyRepository(c *gin.Context, from string, amount int) (domain.Account, error) {
	var res domain.Account
	bankScope := r.bucket.Scope("bank")
	acctCol := bankScope.Collection("accounts")
	userID := strings.ToLower(from)
	userAcct, err := r.GetAccountRepository(c, userID)
	if err != nil {
		log.Println("error getting account")
		return res, err
	}
	balance := userAcct.Balance
	if amount > balance {
		log.Println("lower balance")
		return domain.Account{}, errors.New("insufficient funds")
	}
	newBalance := balance - amount
	// prepare response
	res = domain.Account{
		UserName: userID,
		Balance:  newBalance,
	}
	_, err = acctCol.Upsert(userID, &res, nil)
	if err != nil {
		log.Println("updating balance:", err)
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
	if err := bcrypt.CompareHashAndPassword([]byte(userPass), []byte(password)); err != nil {
		return err
	}
	return nil
}

// RegisterUserRepository :
func (r fintechRepository) RegisterUserRepository(c *gin.Context, userName, email, password string) (string, error) {
	var res domain.User
	bytePass := []byte(password)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	pass := string(hPass)
	res = domain.User{
		UserName: userName,
		Email:    email,
		Password: pass,
		Role:     "member",
	}
	bankScope := r.bucket.Scope("bank")
	userCol := bankScope.Collection("users")
	_, err := userCol.Insert(userName, &res, nil)
	if err != nil {
		return "", err
	}
	acctType := domain.Account{
		UserName: userName,
		Balance:  2000,
	}
	acctCol := bankScope.Collection("accounts")
	_, err = acctCol.Insert(userName, &acctType, nil)
	if err != nil {
		return "", err
	}

	return res.UserName, nil
}

// GetUserRepository GetUserNameRepository :
func (r fintechRepository) GetUserRepository(c *gin.Context, userName string) (domain.UserType, error) {
	var inUser domain.User
	var resp domain.UserType
	bankScope := r.bucket.Scope("bank")
	userCol := bankScope.Collection("users")
	// Get the document back
	getResult, err := userCol.Get(userName, nil)
	if errors.Is(err, gocb.ErrDocumentNotFound) {
		return resp, errors.New("user not found")
	} else if err != nil {
		return resp, err
	}
	err = getResult.Content(&inUser)
	if err != nil {
		return resp, err
	}
	//get user account balance
	accountBalance, err := r.GetAccountRepository(c, userName)
	if err != nil {
		return resp, nil
	}
	//removing username from response
	accountBalance.UserName = ""
	resp = domain.UserType{
		UserName: inUser.UserName,
		Email:    inUser.Email,
		Account:  accountBalance,
	}
	// reset password to empty
	return resp, nil
}

// GetAccountRepository :
func (r fintechRepository) GetAccountRepository(c *gin.Context, userName string) (domain.Account, error) {
	var userAcct domain.Account
	log.Println("username", userName)
	bankScope := r.bucket.Scope("bank")
	col := bankScope.Collection("accounts")
	// Get the document back
	getResult, err := col.Get(userName, nil)
	if errors.Is(err, gocb.ErrDocumentNotFound) {
		return userAcct, errors.New("user not found")
	} else if err != nil {
		return userAcct, err
	}
	err = getResult.Content(&userAcct)
	if err != nil {
		return userAcct, err
	}
	return userAcct, nil
}

// AddToTransaction :
func (r fintechRepository) AddToTransaction(c *gin.Context, from, to string, amount int) error {
	// add transfer to transactions
	txUUID, err := uuid.NewRandom()
	if err != nil {
		log.Println("uuid error:", err)
		return err
	}
	txID := txUUID.String()
	tx := domain.Transaction{
		TransferID: txID,
		From:       from,
		To:         to,
		Amount:     amount,
	}
	bankScope := r.bucket.Scope("bank")
	txCol := bankScope.Collection("transactions")
	_, err = txCol.Insert(txID, &tx, nil)
	if err != nil {
		log.Println("insert error:", err)
		return err
	}
	return nil
}

// NewFintechRepository :
func NewFintechRepository(cluster *gocb.Cluster, bucketName string) (domain.FintechRepository, error) {
	bucket := cluster.Bucket(bucketName)
	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		return nil, err
	}
	return &fintechRepository{cluster: cluster, bucket: bucket}, nil
}
