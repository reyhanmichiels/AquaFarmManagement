package main

import (
	api_call_handler "github.com/reyhanmichiels/AquaFarmManagement/app/api_call/handler"
	api_call_repository "github.com/reyhanmichiels/AquaFarmManagement/app/api_call/repository"
	api_call_usecase "github.com/reyhanmichiels/AquaFarmManagement/app/api_call/usecase"
	farm_handler "github.com/reyhanmichiels/AquaFarmManagement/app/farm/handler"
	farm_repository "github.com/reyhanmichiels/AquaFarmManagement/app/farm/repository"
	farm_usecase "github.com/reyhanmichiels/AquaFarmManagement/app/farm/usecase"
	pond_handler "github.com/reyhanmichiels/AquaFarmManagement/app/pond/handler"
	pond_repository "github.com/reyhanmichiels/AquaFarmManagement/app/pond/repository"
	pond_usecase "github.com/reyhanmichiels/AquaFarmManagement/app/pond/usecase"
	"github.com/reyhanmichiels/AquaFarmManagement/infrastructure"
	"github.com/reyhanmichiels/AquaFarmManagement/infrastructure/database"
	"github.com/reyhanmichiels/AquaFarmManagement/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	//load env
	infrastructure.LoadEnv()

	//create logger
	infrastructure.CreateLogger()

	//connect to database
	database.ConnectToDB()

	//migrate database table
	database.Migrate()

	//init repository
	farmRepository := farm_repository.NewFarmRepository(database.DB)
	pondRepository := pond_repository.NewPondRepository(database.DB)
	apiCallRepository := api_call_repository.NewApiCallRepository(database.DB)

	//init usecase
	farmUsecase := farm_usecase.NewFarmUsecase(farmRepository)
	pondUsecase := pond_usecase.NewPondUsecase(pondRepository, farmRepository)
	apiCallUsecase := api_call_usecase.NewApiCallUsecase(apiCallRepository)

	//init handler
	farmHandler := farm_handler.NewFarmHandler(farmUsecase)
	pondHandler := pond_handler.NewPondHandler(pondUsecase)
	apiCallHandler := api_call_handler.NewApiCallHandler(apiCallUsecase)

	//init rest
	rest := rest.NewRest(gin.New())

	//use middleware
	rest.UseGlobalMiddleware()

	//load route
	rest.HealthCheckRoute()
	rest.FarmRoute(farmHandler)
	rest.PondRoute(pondHandler)
	rest.ApiCallRoute(apiCallHandler)

	//serve app
	rest.Serve()
}
