package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mthcsta/star-wars-api-test/common"
	"github.com/mthcsta/star-wars-api-test/controller"
	"github.com/mthcsta/star-wars-api-test/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	os.Setenv("RUNNING_MAIN", "true")
}

// @contact.name   API Support
// @contact.url    http://localhost:9000
// @contact.email  support@starwarsapi.io
func main() {

	docs.SwaggerInfo.Title = "Star Wars API"
	docs.SwaggerInfo.Description = "Esta é uma API para armazenar informações sobre os planetas do mundo de star wars."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:9000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	if common.Config.ServerEnvironment == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	route := gin.Default()

	apiRoute := route.Group("/api/v1")
	{
		planetController := controller.PlanetController{}
		planets := apiRoute.Group("/planets")
		{
			planets.GET("/", planetController.GetAll)
			planets.POST("/", planetController.Insert)
			planets.DELETE("/:id", planetController.Remove)
		}
		filmController := controller.FilmController{}
		films := apiRoute.Group("/films")
		{
			films.GET("/", filmController.GetAll)
		}
	}

	url := ginSwagger.URL("http://localhost:" + common.Config.ServerPort + "/swagger/doc.json")

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	route.Run(":" + common.Config.ServerPort)

}
