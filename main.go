package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// http POST :1323/any_endpoint 'Authorization:Bearer token'
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		if os.Getenv("API_KEY") != "" {
			return key == os.Getenv("API_KEY"), nil
		} else {
			return key == "testing", nil
		}
	}))

	e.POST("/lunch", setLunch)
	e.POST("/lock", battenDownTheHatches)

	e.Logger.Fatal(e.Start(":1323"))
}

func setLunch(c echo.Context) error {
	if _, err := exec.Command("/home/jlindgren/bin/lunch.rb").Output(); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}

func battenDownTheHatches(c echo.Context) error {
	if _, err := exec.Command(
		"dbus-send",
		"--type=method_call",
		"--dest=org.gnome.ScreenSaver",
		"/org/gnome/ScreenSaver",
		"org.gnome.ScreenSaver.Lock",
	).Output(); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "ok")
}
