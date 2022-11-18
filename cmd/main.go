package main

import (
	"fintech-api/pkg/database"
	fintechhandler "fintech-api/pkg/delivery/http"
	"fintech-api/pkg/repository"
	"fintech-api/pkg/usecase"
	"fintech-api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	config := utils.GetConfigs()

	couchDb, err := database.InitDatabase(config.CouchDbConfig)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	log.Println("successful db and cache connection")
	r := gin.Default()

	fintechRepo, err := repository.NewFintechRepository(couchDb, config.BucketName)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fintechUseCase := usecase.NewFintechUseCase(fintechRepo)

	fintechhandler.NewFintechHandler(r, fintechUseCase)
	err = r.Run(config.Port)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}