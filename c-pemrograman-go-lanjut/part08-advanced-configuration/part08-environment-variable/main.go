package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	configenv "github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := configenv.Load(".env"); err != nil {
		logrus.Error("Cannot load env: ", err.Error())
		return
	}

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Cannot get app.config.json: ", err.Error())
		return
	}

	r := echo.New()
	server := new(http.Server)
	server.Addr = ":" + os.Getenv("SERVER_PORT")
	if serverReadTimeout := os.Getenv("SERVER_READ_TIMEOUT_IN_MINUTE"); serverReadTimeout != "" {
		duration, _ := strconv.Atoi(serverReadTimeout)
		logrus.Info(fmt.Sprintf("Read timeout duration: %d minutes\n", duration))
		server.ReadTimeout = time.Duration(duration) * time.Minute
	}
	if serverWriteTimeout := os.Getenv("SERVER_WRITE_TIMEOUT_IN_MINUTE"); serverWriteTimeout != "" {
		duration, _ := strconv.Atoi(serverWriteTimeout)
		logrus.Info(fmt.Sprintf("Write timeout duration: %d minutes\n", duration))
		server.WriteTimeout = time.Duration(duration) * time.Minute
	}

	r.GET("/", routeIndex)
	r.GET("/foods", routeGetFoods)
	r.GET("/clubs", routeGetClubs)

	r.Logger.Info(fmt.Sprintf("Running %s\n", os.Getenv("APP_NAME")))
	r.Logger.Fatal(r.StartServer(server))
}

var routeIndex = echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	output := make(chan []byte)
	go func() {
		data := []byte("Welcome to my api")
		output <- data
	}()
	w.Write(<-output)
}))

var routeGetFoods = echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type foodModel struct {
		Code  string  `json:"code"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	output := make(chan []foodModel)
	go func() {
		var foods []foodModel
		foodAppConfig := viper.GetStringMap("foods")
		foodInterfaces := foodAppConfig["data"].([]interface{})
		for _, food := range foodInterfaces {
			detail := foodModel{
				Code:  food.(map[string]interface{})["code"].(string),
				Name:  food.(map[string]interface{})["code"].(string),
				Price: food.(map[string]interface{})["price"].(float64),
			}
			foods = append(foods, detail)
		}
		output <- foods
	}()
	outputJSON, _ := json.Marshal(<-output)
	w.Write(outputJSON)
}))

var routeGetClubs = echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type clubModel struct {
		Code  string  `json:"code"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	output := make(chan []clubModel)
	go func() {
		var clubs []clubModel
		clubInterfaces := viper.Get("clubs").([]interface{})
		for _, club := range clubInterfaces {
			detail := clubModel{
				Code:  club.(map[string]interface{})["code"].(string),
				Name:  club.(map[string]interface{})["name"].(string),
				Price: club.(map[string]interface{})["position"].(float64),
			}
			clubs = append(clubs, detail)
		}
		output <- clubs
	}()
	outputJSON, _ := json.Marshal(<-output)
	w.Write(outputJSON)
}))
