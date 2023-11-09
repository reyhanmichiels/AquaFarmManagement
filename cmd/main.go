package main

import (
	"github.com/reyhanmichiels/AquaFarmManagement/infrastructure"
	"github.com/reyhanmichiels/AquaFarmManagement/infrastructure/database"
	"github.com/reyhanmichiels/AquaFarmManagement/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	//load env
	infrastructure.LoadEnv()

	//connect to database
	database.ConnectToDB()

	//init rest
	rest := rest.NewRest(gin.New())

	//load route
	rest.HealthCheckRoute()

	//serve app
	rest.Serve()
}
