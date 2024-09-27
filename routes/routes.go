package routes

import (
	"fmt"
	config "golang-beer-example/configs"
	"golang-beer-example/database"
	router "golang-beer-example/modules/routes"

	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	config.LoadConfig()

	err := database.InitMongoDB()
	if err != nil {
		fmt.Println(err.Error())

	}

	err = database.InitialMariaDatabase()
	if err != nil {
		fmt.Println(err.Error())
	}

	port := config.AppConfig.Service.Port
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if err != nil {
		panic(err)
	} else {
		app := gin.Default()
		api := app.Group("/api")
		router.BeerHandlerRoute(api)
		app.Run(":" + port)
	}
}
