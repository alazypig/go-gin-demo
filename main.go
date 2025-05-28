package main

import (
	"io"
	"os"

	"gilab.com/pragmaticreviews/golang-gin-poc/api"
	"gilab.com/pragmaticreviews/golang-gin-poc/controller"
	"gilab.com/pragmaticreviews/golang-gin-poc/docs"
	"gilab.com/pragmaticreviews/golang-gin-poc/logger"
	"gilab.com/pragmaticreviews/golang-gin-poc/middlewares"
	"gilab.com/pragmaticreviews/golang-gin-poc/repository"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"github.com/gin-gonic/gin"

	// gindupm "github.com/tpkeeper/gin-dump"
	_ "github.com/tpkeeper/gin-dump"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()

	videoService service.VideoService = service.New(videoRepository)
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutPut() {
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	logger.InitLogger()
}

func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Golang Gin POC"
	docs.SwaggerInfo.Description = "Golang Gin POC"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	defer videoRepository.CloseDB()

	setupLogOutPut()

	server := gin.New()
	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(middlewares.Logger(), gin.Recovery())
	// server.Use(gindupm.Dump())

	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
