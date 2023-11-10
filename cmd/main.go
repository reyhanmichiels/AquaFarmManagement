package main

import (
	farm_handler "github.com/reyhanmichiels/AquaFarmManagement/app/farm/handler"
	farm_repository "github.com/reyhanmichiels/AquaFarmManagement/app/farm/repository"
	farm_usecase "github.com/reyhanmichiels/AquaFarmManagement/app/farm/usecase"
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

	//migrate database table
	database.Migrate()

	//init repository
	farmRepository := farm_repository.NewFarmRepository(database.DB)

	//init usecase
	farmUsecase := farm_usecase.NewFarmUsecase(farmRepository)

	//init handler
	farmHandler := farm_handler.NewFarmHandler(farmUsecase)

	//init rest
	rest := rest.NewRest(gin.New())

	//load route
	rest.HealthCheckRoute()
	rest.FarmRoute(farmHandler)

	//serve app
	rest.Serve()
}
