package routers

import (
	"github.com/labstack/echo/v4"
	"ims-intro/handlers"
	"ims-intro/middleware"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.POST("/login", handlers.Login)
	e.POST("/signup", handlers.Signup)

	e.GET("/products", handlers.GetProducts)

	productsGroup := e.Group("/products")
	productsGroup.Use(middleware.AuthAdmin)

	productsGroup.POST("", middleware.AuthAdmin(handlers.CreateProduct))
	productsGroup.PUT("/:id", middleware.AuthAdmin(handlers.UpdateProduct))
	productsGroup.DELETE("/:id", middleware.AuthAdmin(handlers.DeleteProduct))

	return e
}
