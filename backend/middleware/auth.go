package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"ims-intro/models"
	"net/http"
	"strings"
)

var JwtKey = []byte("your_secret_key")

func AuthAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")
		if tokenStr == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "No token provided")
		}

		tokenStr = strings.Split(tokenStr, "Bearer ")[1]
		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil || !token.Valid || claims.Role != "admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not authorized")
		}

		c.Set("user", claims.Username)
		return next(c)
	}
}
