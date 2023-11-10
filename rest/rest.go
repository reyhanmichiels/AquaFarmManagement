package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	farm_handler "github.com/reyhanmichiels/AquaFarmManagement/app/farm/handler"
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
			"status": "success",
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

func (rest *Rest) Serve() {
	rest.engine.Run()
}
