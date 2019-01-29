package main

import (
	"Mock-API-Data/api/router"
	"Mock-API-Data/config"
	"Mock-API-Data/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
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

	gin.SetMode(gin.DebugMode)

	apiRouter := router.InitDashboardRouter(storage)

	mockRouter := router.InitMockRouter(storage)

	fmt.Println("start-up success ....")
	go func() {
		addr := fmt.Sprintf("%s:%v", conf.DashboardAddr, conf.DashboardPort)
		fmt.Println("Listen Dashboard: ", addr)
		if err := apiRouter.Run(addr); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}()

	go func() {
		addr := fmt.Sprintf("%s:%v", conf.MockAddr, conf.MockPort)
		fmt.Println("Listen Mock: ", addr)
		if err := mockRouter.Run(addr); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}()

	// 阻塞
	select {}
}
