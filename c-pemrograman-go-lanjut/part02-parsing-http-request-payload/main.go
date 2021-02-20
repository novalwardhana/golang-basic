package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type user struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type food struct {
	Name  string `json:"name" form:"name" query:"name"`
	Price int64  `json:"price" form:"price" query:"price"`
}

func main() {
	r := echo.New()
	r.Any("/user", routeUser)
	r.Any("/food", routeFood)

	r.Start(":9000")
}

func routeUser(c echo.Context) error {
	data := new(user)
	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func routeFood(c echo.Context) error {
	data := new(food)
	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}
