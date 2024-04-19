package main

import (
	"github.com/ericsison/bill-inquiry/config"
	"github.com/ericsison/bill-inquiry/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	config.LoadEnv()

	config.InitDB()

	defer config.DestroyDB(config.GetDBConnection())

	gin.SetMode(os.Getenv("GIN_MODE"))

	app := gin.Default()

	if err := app.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	routes.RegisterConsumerRoutes(app)
	routes.RegisterConsumerWithArrears(app)

	err := app.Run(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))

	if err != nil {
		log.Fatal(err)
	}
}
