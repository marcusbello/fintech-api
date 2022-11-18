package utils

type config struct {
	CouchDbConfig
	BucketName string
	Port       string
}

type CouchDbConfig struct {
	Host     string
	User     string
	Password string
}

func GetConfigs() *config {
	couchConfig := CouchDbConfig{
		Host:     "localhost",
		User:     "Administrator",
		Password: "couchbase",
	}
	bucketName := "fintech"

	port := ":3030"
	return &config{
		CouchDbConfig: couchConfig,
		BucketName:    bucketName,
		Port:          port,
	}
}
