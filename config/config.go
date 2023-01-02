package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type config struct {
	CouchBaseConfig
	CouchBaseBucket string
	HTTPPort        string
}

type CouchBaseConfig struct {
	Host     string
	User     string
	Password string
}

func GetConfigs() *config {
	getCredentials := &config{}
	//couchbase config
	getCredentials.CouchBaseConfig = getCouchBase()
	//couchbase bucket
	getCredentials.CouchBaseBucket = getBucket()

	port := getPort() //getting port from env
	getCredentials.HTTPPort = port
	return getCredentials
}

// getCouchBase is the helper function that gets couchbase credentials from the environment variable
func getCouchBase() CouchBaseConfig {
	DbConfig := CouchBaseConfig{}
	host := os.Getenv("APP_COUCHBASE_HOST")
	user := os.Getenv("APP_COUCHBASE_USER")
	password := os.Getenv("APP_COUCHBASE_PASSWORD")

	if host == "" || user == "" || password == "" {
		log.Println("Empty couchbase credentials in environment variable, using default credentials")
		DefaultConfig := CouchBaseConfig{
			Host:     "localhost",
			User:     "Administrator",
			Password: "couchbase",
		}
		return DefaultConfig
	}
	DbConfig.Host = host
	DbConfig.User = user
	DbConfig.Password = password

	return DbConfig
}

// getBucket is the helper function that gets couchbase bucket name from the environment variable
func getBucket() string {
	bucketName := os.Getenv("APP_COUCHBASE_BUCKET")
	if bucketName == "" {
		log.Println("Couchbase bucket name not in environment variable, using default bucket name (fintech)")
		return "fintech"
	}

	return bucketName
}

// getPort is the helper function that gets the http port from the environment variable
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Port number not in environment variable, using default port (:3030)")
		return ":3030"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Failed to parse PORT environment variable")
	}

	return fmt.Sprintf(":%v", portInt)
}
