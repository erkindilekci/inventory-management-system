package controller

import (
	"github.com/labstack/echo/v4"
	"ims-intro/pkg/controller/request"
	"ims-intro/pkg/controller/response"
	"ims-intro/pkg/domain"
	"ims-intro/pkg/middleware"
	"ims-intro/pkg/service"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{productService}
}

func (controller *ProductController) RegisterProductRoutes(e *echo.Echo) {
	productsGroup := e.Group("/products")
	productsGroup.Use(middleware.AuthMiddleware)

	productsGroup.GET("", controller.GetAllProducts)
	productsGroup.POST("", controller.AddNewProduct)
	productsGroup.PUT("/:id", controller.UpdateProductById)
	productsGroup.DELETE("/:id", controller.DeleteProductById)
}

func (controller *ProductController) GetAllProducts(c echo.Context) error {
	category := c.QueryParam("category")
	var products []*domain.Product

	if len(category) == 0 {
		products = controller.productService.GetAllProducts()
	} else {
		products = controller.productService.GetAllProductsByCategory(category)
	}

	return c.JSON(http.StatusOK, response.ToProductResponseList(products))
}

func (controller *ProductController) AddNewProduct(c echo.Context) error {
	addProductResponse := new(request.AddProductRequest)

	err := c.Bind(addProductResponse)
	if err != nil || addProductResponse == nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: unable to bind the provided data to the product structure"))
	}

	err = controller.productService.Add(addProductResponse.ToModel())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusCreated)
}

func (controller *ProductController) UpdateProductById(c echo.Context) error {
	param := c.Param("id")
	if param == "" {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: no product id specified"))
	}

	productId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: product id must be an integer"))
	}

	addProductResponse := new(request.AddProductRequest)
	err = c.Bind(addProductResponse)
	if err != nil || addProductResponse == nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: unable to bind the provided data to the product structure"))
	}

	err = controller.productService.UpdateProductById(addProductResponse.ToModel(), int64(productId))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusOK)
}

func (controller *ProductController) DeleteProductById(c echo.Context) error {
	param := c.Param("id")
	if param == "" {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: no product id specified"))
	}

	productId, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid request: product id must be an integer"))
	}

	err = controller.productService.DeleteById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusOK)
}
