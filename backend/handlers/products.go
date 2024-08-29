package handlers

import (
	"github.com/labstack/echo/v4"
	"ims-intro/models"
	"net/http"
	"strconv"
)

func GetProducts(c echo.Context) error {
	rows, err := models.DB.Query("SELECT id, name, price, quantity, category FROM products")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity, &p.Category); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		products = append(products, p)
	}

	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := models.DB.QueryRow("INSERT INTO products (name, price, quantity, category) VALUES ($1, $2, $3, $4) RETURNING id",
		product.Name, product.Price, product.Quantity, product.Category).Scan(&product.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err = models.DB.Exec("UPDATE products SET name = $1, price = $2, quantity = $3, category = $4 WHERE id = $5",
		product.Name, product.Price, product.Quantity, product.Category, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err = models.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
