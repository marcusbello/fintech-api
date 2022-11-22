package config

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
	//couchbase config
	couchConfig := CouchBaseConfig{
		Host:     "localhost",
		User:     "Administrator",
		Password: "couchbase",
	}
	bucketName := "fintech"

	port := ":3030"
	return &config{
		CouchBaseConfig: couchConfig,
		CouchBaseBucket: bucketName,
		HTTPPort:        port,
	}
}
