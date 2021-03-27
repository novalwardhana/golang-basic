package main

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	r := echo.New()

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	viper.WatchConfig()
	viper.OnConfigChange(func(f fsnotify.Event) {
		logrus.Info(fmt.Sprintf("Config file changed: %s\n", f.Name))
	})

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("An error occured: ", err.Error())
		return
	}

	r.GET("/", routeIndex)
	r.GET("/data", routeGetData)

	r.Logger.Info("Start app %s", viper.GetString("appName"))
	r.Logger.Fatal(r.Start(":" + viper.GetString("server.port")))
}

func routeIndex(c echo.Context) error {
	output := make(chan string)
	go func() {
		data := "Welcome to my api"
		output <- data
	}()
	return c.JSON(http.StatusOK, map[string]interface{}{"message": <-output})
}

func routeGetData(c echo.Context) error {
	output := make(chan map[string]interface{})
	go func() {
		output <- viper.GetStringMap("data")
	}()
	return c.JSON(http.StatusOK, <-output)
}
