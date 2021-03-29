package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var hashCookie = securecookie.New([]byte("new-cookie"), []byte("1a2b3b4d5e6f7g8h"))

type M map[string]interface{}

var cookieIndex = "index"
var cookieProfile = "profile"

func setCookie(c echo.Context, name string, data M) error {
	encodedCookie, err := hashCookie.Encode(name, data)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     name,
		Value:    encodedCookie,
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
	}
	http.SetCookie(c.Response(), cookie)
	return nil
}

func getCookie(c echo.Context, name string) (M, error) {
	cookie, err := c.Request().Cookie(name)
	if err != nil {
		return nil, err
	}
	data := M{}
	if err := hashCookie.Decode(name, cookie.Value, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func removeCookie(c echo.Context, name string) {
	cookie := &http.Cookie{}
	cookie.Name = name
	cookie.Path = "/"
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(0, 0)
	http.SetCookie(c.Response(), cookie)
}

func main() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("An error occured: %s", err.Error())
		return
	}

	appName := viper.GetString("app_name")
	serverConfig := viper.GetStringMap("server")
	logrus.Info(fmt.Sprintf("Start %s \n...", appName))

	r := echo.New()
	server := new(http.Server)
	server.Addr = fmt.Sprintf(":%d", int(serverConfig["port"].(float64)))
	server.WriteTimeout = time.Duration(int(serverConfig["write_timeout"].(float64))) * time.Second
	server.ReadTimeout = time.Duration(int(serverConfig["read_timeout"].(float64))) * time.Second

	r.GET("/", routeIndex)
	r.GET("/profile", routeGetProfile)
	r.GET("/remove-cookie", routeRemoveCookie)
	r.GET("/remove-cookie-profile", routeRemoveCookieProfile)

	r.StartServer(server)
}

func routeIndex(c echo.Context) error {
	data, err := getCookie(c, cookieIndex)
	if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
		return err
	}
	if data == nil {
		newData := M{"name": cookieIndex, "ID": 20210329}
		err := setCookie(c, cookieIndex, newData)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, newData)
	}
	return c.JSON(http.StatusOK, data)
}

func routeGetProfile(c echo.Context) error {
	data, err := getCookie(c, cookieProfile)
	if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
		return err
	}
	if data == nil {
		newData := M{"type": cookieProfile, "name": "Noval Wardhana", "github_profile": "github.com/novalwardhana"}
		err := setCookie(c, cookieProfile, newData)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, newData)
	}

	return c.JSON(http.StatusOK, data)
}

func routeRemoveCookie(c echo.Context) error {
	removeCookie(c, cookieIndex)
	/* return c.JSON(http.StatusOK, map[string]interface{}{"message": "Cookie have been deleted"}) */
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func routeRemoveCookieProfile(c echo.Context) error {
	removeCookie(c, cookieProfile)
	/* return c.JSON(http.StatusOK, map[string]interface{}{"message": "Cookie have been deleted"}) */
	return c.Redirect(http.StatusTemporaryRedirect, "/profile")
}
