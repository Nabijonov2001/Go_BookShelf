package api

import (
	_ "github.com/abdukhashimov/golang-hex-architecture/api/docs"
	v1 "github.com/abdukhashimov/golang-hex-architecture/api/handlers/v1"
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/pkg/middleware"
	"github.com/abdukhashimov/golang-hex-architecture/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type RouterOptions struct {
	Cfg     *config.Config
	Log     *zap.Logger
	Service service.ServiceI
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.NewHandler(&v1.HandlerOptions{
		Cfg:     opt.Cfg,
		Log:     opt.Log,
		Service: opt.Service,
	})

	apiV1 := router.Group("/v1")

	// user routes
	{
		apiV1.POST("/signup", handlerV1.CreateUser)
		apiV1.GET("/myself", middleware.Auth, handlerV1.GetOneUser)
	}

	// book routes
	{
		apiV1.GET("/books", middleware.Auth, handlerV1.GetAllBooks)
		apiV1.POST("/books", middleware.Auth, handlerV1.CreateBook)
		apiV1.PATCH("/books/:id", middleware.Auth, handlerV1.UpdateBook)
		apiV1.DELETE("/books/:id", middleware.Auth, handlerV1.DeleteBook)
	}

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
