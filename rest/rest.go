package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api_call_handler "github.com/reyhanmichiels/AquaFarmManagement/app/api_call/handler"
	farm_handler "github.com/reyhanmichiels/AquaFarmManagement/app/farm/handler"
	pond_handler "github.com/reyhanmichiels/AquaFarmManagement/app/pond/handler"
	"github.com/reyhanmichiels/AquaFarmManagement/middleware"
)

type Rest struct {
	engine *gin.Engine
}

func NewRest(engine *gin.Engine) Rest {
	return Rest{
		engine: engine,
	}
}

func (rest *Rest) HealthCheckRoute() {
	rest.engine.GET("/api/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "successfully run health check",
		})
	})
}

func (rest *Rest) FarmRoute(farmHandler *farm_handler.FarmHandler) {
	rest.engine.GET("/api/farms", farmHandler.Get)
	rest.engine.POST("/api/farms", farmHandler.Create)
	rest.engine.GET("/api/farms/:farmId", farmHandler.GetFarmById)
	rest.engine.PUT("/api/farms/:farmId", farmHandler.Update)
	rest.engine.DELETE("/api/farms/:farmId", farmHandler.Delete)
}

func (rest *Rest) PondRoute(pondHanler *pond_handler.PondHandler) {
	rest.engine.GET("/api/ponds", pondHanler.Get)
	rest.engine.POST("/api/ponds", pondHanler.Create)
	rest.engine.GET("/api/ponds/:pondId", pondHanler.GetPondById)
	rest.engine.PUT("/api/ponds/:pondId", pondHanler.Update)
	rest.engine.DELETE("/api/ponds/:pondId", pondHanler.Delete)
}

func (rest *Rest) ApiCallRoute(apiCallHandler *api_call_handler.ApiCallHandler) {
	rest.engine.GET("/api/api-calls", apiCallHandler.Get)

}

func (rest *Rest) UseGlobalMiddleware() {
	rest.engine.Use(middleware.LogEvent)
	rest.engine.Use(middleware.RecordApiCallMiddleware)
}

func (rest *Rest) Serve() {
	rest.engine.Run()
}
