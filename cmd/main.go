package main

import (
	getconfig "fintech-api/config"
	"fintech-api/pkg/database/couchbase"
	fintechhandler "fintech-api/pkg/delivery/http"
	"fintech-api/pkg/usecase"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	config := getconfig.GetConfigs()

	couchDb, err := couchbase.InitCouchBase(config.CouchBaseConfig)
	if err != nil {
		log.Fatalf("failed to establish db connection %v\n", err)
	}

	log.Println("Couchbase handshake successful")

	//
	r := gin.Default()

	//fintechRepo, err := repository.NewFintechRepository(couchDb, config.CouchBaseBucket)
	//if err != nil {
	//	log.Fatalf("main.fintechRepo: Couchbase bucket or server error: %v\n", err)
	//}
	fintechRepo, err := couchbase.NewCouchbaseRepo(couchDb, config.CouchBaseBucket)
	if err != nil {
		log.Fatalf("main.fintechRepo: Couchbase bucket or server error: %v\n", err)
	}
	fintechUseCase := usecase.NewFintechUseCase(fintechRepo)

	fintechhandler.NewFintechHandler(r, fintechUseCase)
	err = r.Run(config.HTTPPort)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
