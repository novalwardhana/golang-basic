package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type customValidation struct {
	Validation *validator.Validate
}

func (cv *customValidation) Validate(i interface{}) error {
	return cv.Validation.Struct(i)
}

func main() {
	r := echo.New()
	r.Validator = &customValidation{Validation: validator.New()}

	r.Any("/user", routeUser)
	r.POST("/food", routeFood)

	r.Logger.Error(r.Start(":9000"))
}

type user struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int64  `json:"age" validate:"gte=15,lte=50"`
}

func routeUser(c echo.Context) error {
	data := new(user)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}
	if err := c.Validate(data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "message": "Data valid", "data": data})
}

type food struct {
	Code  string `json:"code" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Price int64  `json:"price" validate:"gte=1000,lte=100000"`
}

func routeFood(c echo.Context) error {
	data := new(food)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}
	if err := c.Validate(data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "message": "Data valid", "data": data})
}
