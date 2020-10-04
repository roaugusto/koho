package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
	ctrl "github.com/roaugusto/kohobalance/internal/records/controllers"
	"go.mongodb.org/mongo-driver/mongo"

	_ "github.com/roaugusto/kohobalance/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//Routes Set routes for the application
func Routes(e *echo.Echo, col *mongo.Collection, host string, port string) {

	// Routes for swagger
	urlDocSwagger := fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)
	url := echoSwagger.URL(urlDocSwagger)
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))

	// Routes for records
	records := ctrl.RecordHandler{Repo: col}
	e.POST("/api/funds", records.CreateRecordsFromFile)
	e.POST("/api/funds-write-result-db", records.CreateRecordsFromFileDb)
	e.POST("/api/funds-body-req", records.CreateRecordsBodyRequest)

	e.GET("/api/funds/download", records.GetFile)
	e.GET("/api/funds/result", records.GetRecordsFromDB)

}
