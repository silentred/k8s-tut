package main

import (
	"fmt"

	"math/rand"

	"github.com/labstack/echo"
	"github.com/silentred/kassadin"
	"github.com/silentred/kassadin/filter"
)

var id int

func main() {
	id = rand.Int()
	app := kassadin.NewApp()
	app.RegisterRouteHook(route)
	app.Start()
}

func route(app *kassadin.App) error {
	h := filter.GetPrometheusLogHandler()
	app.Route.GET("/metrics", h)

	app.Route.GET("/hello*", helloWorld, filter.Recover(), filter.Logger(app.DefaultLogger()), filter.Metrics())
	return nil
}

func helloWorld(ctx echo.Context) error {
	return ctx.String(200, fmt.Sprintf("ID:%d path: %s", id, ctx.Request().URL.Path))
}
