package main

import (
	"Mock-API-Data/api/router"
	"Mock-API-Data/config"
	"Mock-API-Data/storage"
	"fmt"
	"log"
)

func main() {

	confPath := "./conf.yaml"
	conf, err := config.Load(confPath)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := storage.NewStorage(conf)
	if err != nil {
		log.Fatal(err)
	}

	apiRouter := router.InitRouter(storage)

	fmt.Println("start-up success ....")

	if err := apiRouter.Run(fmt.Sprintf("%s:%v", conf.DashboardAddr, conf.DashboardPort)); err != nil {
		fmt.Println(err)
	}
}
