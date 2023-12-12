package main

import (
	"fmt"
	"github.com/YuStep/wbl0/config"
	_ "github.com/YuStep/wbl0/docs"
	"github.com/YuStep/wbl0/internal/handler"
	"github.com/YuStep/wbl0/internal/repository"
	"github.com/YuStep/wbl0/internal/service"
	"github.com/YuStep/wbl0/pkg/cache"
	"github.com/YuStep/wbl0/pkg/nats"
	_ "github.com/lib/pq"
	"log"
	"time"
)

//	@title			Go WB API
//	@version		1.0
//	@description	This is a simple RESTful Service API written in Go using Gin web framework, PostgreSQL database,NATS queue with in-memory caching.

// @contact.name   API Support
// @contact.url    https://github.com/YuStep/
// @contact.email  support@mail.io

//	@license.name  MIT
//	@license.url   https://github.com/YuStep/wbl0/blob/main/LICENSE

// @host      localhost:8080
// @BasePath  /v1

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Error at starting: %v", err)
	}

	ns := nats.NewNats(&cfg.Nats)
	if ns == nil {
		log.Fatalf("Error creating NATS connection")
	}
	defer ns.Close()
	fmt.Println("Nats server started, connection registered")

	postgresConnection, err := repository.Connect(&cfg.PG)

	if err != nil {
		fmt.Printf("Error at Postgre connection: %v", err)
	}
	defer postgresConnection.Close()

	orderCache := cache.NewCache()

	repos := repository.NewRepository(postgresConnection, orderCache)
	dbCreationError := repos.Order.CreateTable()

	if dbCreationError != nil {
		fmt.Printf("Error at table creating: %v", dbCreationError)
	}

	services := service.NewService(repos)
	services.Preload()

	go func() {
		for {
			order := repository.GenerateOrder()
			fmt.Println("Order sent")
			err := ns.Publish(*order)

			if err != nil {
				fmt.Printf("Error at publishing: %v\n", err)
			}
			time.Sleep(30 * time.Second)
		}
	}()

	go func() {
		for {
			mes, err := ns.Subscribe()
			fmt.Println("Order received")
			if err != nil {
				fmt.Printf("Error at subscribing: %v", err)
				fmt.Printf("Error at Unmarshaling: %v", err)
				time.Sleep(30 * time.Second)
			} else {
				services.SaveOrder(*mes)
			}
		}
	}()

	handlers := handler.NewHandler(services)

	serv := handlers.InitRoutes()

	err = serv.Run(":8000")
	if err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}
