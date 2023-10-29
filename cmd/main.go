package main

import (
	handlers "github.com/romendoza-meli/code-review/cmd/handler"
	httpErr "github.com/romendoza-meli/code-review/cmd/http-error"
	"github.com/romendoza-meli/code-review/internal/mapper"
	"github.com/romendoza-meli/code-review/internal/vehicle/loader"
	"github.com/romendoza-meli/code-review/internal/vehicle/repository"
	"github.com/romendoza-meli/code-review/internal/vehicle/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	godotenv.Load(".env")

	// dependencies
	ldVh := loader.NewLoaderVehicleJSON(os.Getenv("FILE_PATH_VEHICLES_JSON"))
	dbVh, err := ldVh.Load()
	if err != nil {
		panic(err)
	}

	rpVh := repository.NewRepositoryVehicleInMemory(dbVh)
	svVh := service.NewServiceVehicleDefault(rpVh, service.ErrorAdapter)
	ctVh := handlers.NewControllerVehicle(svVh, httpErr.ErrorAdapter, mapper.NewStructMapper())

	// server
	rt := gin.New()
	// -> middlewares
	rt.Use(gin.Recovery())
	rt.Use(gin.Logger())
	// -> handlers
	api := rt.Group("/api/v1")
	grVh := api.Group("/vehicles")
	grVh.GET("", ctVh.GetAll())
	grVh.PATCH("/:id/update_fuel", ctVh.PatchFuel())
	grVh.PUT("/:id/update_fuel", ctVh.PutFuel())
	grVh.POST("/batch", ctVh.Batch())
	grVh.POST("/post", ctVh.Post())
	grVh.DELETE("/:id", ctVh.Delete())
	grVh.GET("/color/:color/year/:year", ctVh.GetByColorAndYear())
	grVh.GET("/average_capacity/brand/:brand", ctVh.GetAverageCapacityByBrand())
	grVh.GET("/dimensions", ctVh.GetByDimensions())
	grVh.GET("/weight", ctVh.GetByWeight())
	grVh.GET("/transmission/:type", ctVh.GetByTransmission())

	// run
	if err := rt.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}
}
