package main

import (
	"log"
	"os"

	database "github.com/mauroccvieira/restapi-echo-go/db"
	_ "github.com/mauroccvieira/restapi-echo-go/docs"
	"github.com/mauroccvieira/restapi-echo-go/handlers"
	"github.com/mauroccvieira/restapi-echo-go/logger"
	"github.com/mauroccvieira/restapi-echo-go/services"
	"github.com/mauroccvieira/restapi-echo-go/stores"
	"go.uber.org/zap"
)

var GO_ENV = os.Getenv("GO_ENV")

// @title Go RestAPI API v1
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8888
// @BasePath /
// @schemes http
func main() {
	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(GO_ENV == "development")
	if err != nil {
		logger.Fatal("failed to connect to the database", zap.Error(err))
	}
	defer db.Close()

	e := handlers.Echo()

	s := stores.New(db)
	ss := services.New(s)
	h := handlers.New(ss)

	handlers.SetDefault(e)
	handlers.SetApi(e, h, nil)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8888"
	}

	e.Start(":" + PORT)
}
