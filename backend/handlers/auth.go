package handlers

import (
	"database/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"ims-intro/middleware"
	"ims-intro/models"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	storedUser := models.User{}
	err := models.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = $1", user.Username).Scan(&storedUser.ID, &storedUser.Username, &storedUser.Password, &storedUser.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, "User not found")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Username: storedUser.Username,
		Role:     storedUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = expirationTime
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func Signup(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(hashedPassword)

	err = models.DB.QueryRow("INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id",
		user.Username, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
