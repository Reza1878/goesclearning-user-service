package routes

import (
	"fmt"
	"log"
	"strings"

	productHandlers "github.com/Reza1878/goesclearning/user-service/handler/product"
	handlers "github.com/Reza1878/goesclearning/user-service/handler/user"
	"github.com/Reza1878/goesclearning/user-service/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Routes struct {
	Router  *gin.Engine
	User    *handlers.Handler
	Product *productHandlers.Handler
}

func (r *Routes) SetupRoutes() {
	r.Router = gin.New()
	r.Router.Use(middlewares.EnabledCORS(), middlewares.Logger(r.Router))

	r.setupAPIRoutes()
}

func (r *Routes) setupAPIRoutes() {
	baseURL := viper.GetString("BASE_URL_PATH")
	if baseURL == "" || baseURL == "/" {
		baseURL = "/"
	} else {
		baseURL = "/" + strings.TrimPrefix(baseURL, "/")
	}

	apiGroup := r.Router.Group(baseURL)
	r.configureUserRoutes(apiGroup)
	r.configureProductRoutes(apiGroup)
}

func (r *Routes) configureUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	userGroup.POST("/register", r.User.HandleUserRegister)
	userGroup.POST("/login", r.User.HandleUserLogin)
}

func (r *Routes) configureProductRoutes(router *gin.RouterGroup) {
	productGroup := router.Group("/product")
	productGroup.POST("/", r.Product.InsertProduct)
	productGroup.GET("/", r.Product.ListProduct)
}

func (r *Routes) Run(port string) {
	if r.Router == nil {
		panic("[ROUTER ERROR] Gin Engine has not been initialized. Make sure to call SetupRouter() before Run().")
	}

	addr := fmt.Sprintf(":%s", port)
	if err := r.Router.Run(addr); err != nil {
		panic(fmt.Sprintf("[SERVER ERROR] Failed to start the server on port %s: %v", port, err))
	}

	log.Default().Printf("[INFO] Server running at port: %s", addr)
}
