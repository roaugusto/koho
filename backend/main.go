package main

import (
	"context"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/gommon/log"
	config "github.com/roaugusto/kohobalance/config"
	logger "github.com/roaugusto/kohobalance/internal/log"
	routes "github.com/roaugusto/kohobalance/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientDB   *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
	cfg        config.Properties
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration cannot be read: %v ", err)
	}

	connectURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	clientDB, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Enable to connect to database: %v ", err)
	}

	db = clientDB.Database(cfg.DBName)
	collection = db.Collection(cfg.CollectionName)

}

// @title Swagger Koho Load Funds
// @version 1.0
// @description API REST to load funds based on a file with load amounts.

// @contact.name Rodrigo Santos
// @contact.email ro.augusto@gmail.com

// @host localhost:3333
// @BasePath /
func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType,
			echo.HeaderAccept, "Cache-Control"},
	}))

	logger.Config(e)

	routes.Routes(e, collection, cfg.Host, cfg.Port)

	e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	e.Logger.Infof("Listening on %s:%s", cfg.Host, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

}
