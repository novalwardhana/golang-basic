package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type customValidation struct {
	validate *validator.Validate
}

func (cv *customValidation) Validate(i interface{}) error {
	return cv.validate.Struct(i)
}

const errHandlingType = "advance"

func main() {
	r := echo.New()
	r.Validator = &customValidation{validate: validator.New()}

	r.POST("/user", routeUser)
	r.POST("/food", routeFood)

	if errHandlingType == "basic" {

		r.HTTPErrorHandler = func(err error, c echo.Context) {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			c.Logger().Error(report)
			c.JSON(report.Code, report)
		}

	} else if errHandlingType == "advance" {

		r.HTTPErrorHandler = func(err error, c echo.Context) {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			if castedObject, ok := err.(validator.ValidationErrors); ok {
				for _, err := range castedObject {
					switch err.Tag() {
					case "required":
						report.Message = fmt.Sprintf("%s is required", err.Field())
					case "email":
						report.Message = fmt.Sprintf("%s is not valid", err.Field())
					case "gte":
						report.Message = fmt.Sprintf("%s must greater than %s", err.Field(), err.Param())
					case "lte":
						report.Message = fmt.Sprintf("%s must be lower than %s", err.Field(), err.Param())
					}
				}
			}
			c.Logger().Error(report)
			c.JSON(report.Code, report)
		}

	} else {

		r.HTTPErrorHandler = func(err error, c echo.Context) {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			errPage := fmt.Sprintf("%d.html", http.StatusForbidden)
			if err := c.File(errPage); err != nil {
				c.JSON(report.Code, "An error occured")
			}
		}

	}

	r.Logger.Error(r.Start(":9000"))
}

type user struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int64  `json:"age" validate:"gte=10,lte=30"`
}

func routeUser(c echo.Context) error {
	data := new(user)
	if err := c.Bind(data); err != nil {
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
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
		return err
	}
	if err := c.Validate(data); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"code": http.StatusOK, "message": "Data valid", "data": data})
}
