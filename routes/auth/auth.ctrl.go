package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Register : Register Router
func (AuthRouter) Register(c echo.Context) error {
	type RequestBody struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": "token",
		"user":  "user",
	})
}

// Login : Login Router
func (AuthRouter) Login(c echo.Context) error {
	type RequestBody struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": "token",
		"user":  "user",
	})
}

// Logout : Logout Router
func (AuthRouter) Logout(c echo.Context) error {
	tokenCookie, _ := c.Get("tokenCookie").(*http.Cookie)

	tokenCookie.Value = ""
	tokenCookie.Expires = time.Unix(0, 0)

	c.SetCookie(tokenCookie)

	return c.NoContent(200)
}
