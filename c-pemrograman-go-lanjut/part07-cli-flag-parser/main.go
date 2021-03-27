package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)

/* Call: ./main-arg-parser "Noval APP" 9000
var (
	app        = kingpin.New("App", "Simple app")
	aggAppName = app.Arg("name", "Application Name").Required().String()
	aggAppPort = app.Arg("port", "Application Port").Default("9000").Int()
)
*/

/* Call: ./main --name "Noval" -p 9011*/
var (
	app        = kingpin.New("App", "Simple App")
	aggAppName = app.Flag("name", "Application name").Required().String()
	aggAppPort = app.Flag("port", "Application port").Short('p').Default("9000").Int()
)

var (
	commandAdd             = app.Command("add", "Add new user")
	commandAddFlagOverride = commandAdd.Flag("override", "Override existing user").Short('o').Bool()
	commandAddArgUser      = commandAdd.Arg("user", "username").Required().String()
)

var (
	commandUpdate           = app.Command("update", "Update existing user")
	commandUpdateArgOldUser = commandUpdate.Arg("old", "old username").Required().String()
	commandUpdateArgNewUser = commandUpdate.Arg("new", "new username").Required().String()
)

var (
	commandDelete          = app.Command("delete", "delete user")
	commandDeleteFlagForce = commandDelete.Flag("force", "delete user").Short('f').Bool()
	commandDeleteArgUser   = commandDelete.Arg("user", "delete user").Required().String()
)

func main() {
	/* Call: ./main --name "NovalApp" -p 9011 add --override "Noval" */
	commandAdd.Action(func(ctx *kingpin.ParseContext) error {
		user := *commandAddArgUser
		override := *commandAddFlagOverride
		fmt.Printf("Adding user %s override %t\n", user, override)
		return nil
	})

	/* Call: ./main --name "NovalApp" -p 9011 update "Noval" "Wardhana" */
	commandUpdate.Action(func(ctx *kingpin.ParseContext) error {
		oldUser := *commandUpdateArgOldUser
		newUser := *commandUpdateArgNewUser
		fmt.Printf("Update user from %s to %s\n", oldUser, newUser)
		return nil
	})

	/* Call: ./main --name "NovalAPP" -p 9011 delete --force "Wardhana" */
	commandDelete.Action(func(ctx *kingpin.ParseContext) error {
		user := *commandDeleteArgUser
		flagForce := *commandDeleteFlagForce
		fmt.Printf("Delete user %s force %t\n", user, flagForce)
		return nil
	})

	commandInString := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch commandInString {
	case commandAdd.FullCommand():
		user := *commandAddArgUser
		override := *commandAddFlagOverride
		fmt.Printf("Adding user %s override %t\n", user, override)
	case commandUpdate.FullCommand():
		oldUser := *commandUpdateArgOldUser
		newUser := *commandUpdateArgNewUser
		fmt.Printf("Change user from %s to %s\n", oldUser, newUser)
	case commandDelete.FullCommand():
		user := *commandDeleteArgUser
		flagForce := *commandDeleteFlagForce
		fmt.Printf("Delete user %s force %t\n", user, flagForce)
	}

	appName := *aggAppName
	appPort := fmt.Sprintf(":%d", *aggAppPort)

	fmt.Printf("Start %s in %s", appName, appPort)

	r := echo.New()

	r.GET("/", routeIndex)

	r.Start(appPort)
}

func routeIndex(c echo.Context) error {
	output := make(chan interface{})
	go func() {
		data := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusOK,
			Message: "Welcome to my API...",
		}
		output <- data
	}()
	return c.JSON(http.StatusOK, <-output)
}
