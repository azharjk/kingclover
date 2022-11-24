package app

import (
	"github.com/gin-gonic/gin"
	"kingclover.com/api/pkg/handler"
	"kingclover.com/api/pkg/handler/product"
)

// "/api"
func setupAuthRoutes(r *gin.RouterGroup) {
	r.POST("/register", handler.RegisterHandler)
	r.POST("/login", handler.LoginHandler)
	r.POST("/logout", handler.LogoutHandler)
	r.POST("/token", handler.TokenHandler)
}

// "/api/products"
func setupProductRoutes(r *gin.RouterGroup) {
	r.GET("/", product.ListHandler)
	r.GET("/:id", product.DetailHandler)
	r.POST("/", product.CreateHandler)
	r.PUT("/:id", product.UpdateHandler)
	r.DELETE("/:id", product.DeleteHandler)
}

// "/api/user"
func setupUserRoutes(r *gin.RouterGroup) {
	r.GET("/", handler.UserInfoHandler)
}

func SetupRoutes(r *gin.Engine) {
	r.GET("/", handler.GreetHandler)

	api := r.Group("/api")

	setupAuthRoutes(api)

	user := api.Group("/user")
	setupUserRoutes(user)

	products := api.Group("/products")
	setupProductRoutes(products)
}
