package router

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	log "github.com/tOnkowzl/libs/logx"
	"github.com/tOnkowzl/libs/middleware"
	"github.com/tinnagorn/my-golang-service-template/cachemanager"
	"github.com/tinnagorn/my-golang-service-template/database"
)

var (
	buildstamp = time.Now().String()
	githash    = "developing"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	m := middleware.New(viper.GetString("app.name"))
	m.Skipper = func(c echo.Context) bool {
		return c.Path() == "/builds" || c.Path() == "/health"
	}

	e.Use(m.Build(buildstamp, githash))
	e.Use(m.RequestID())
	e.Use(m.Recover())
	e.Use(m.LogRequestInfo())
	e.Use(m.Logger())

	return e
}

func Runs(router *echo.Echo) {
	log.Infof("starting %s", viper.GetString("app.name"))
	log.Infof("application serve at port %s", viper.GetString("app.port"))
	log.Info(router.Start(":" + viper.GetString("app.port")))
}

func Shutdown(router *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	defer database.Close()
	defer cachemanager.Close()
	if err := router.Shutdown(context.Background()); err != nil {
		log.Fatal("shutdown server: ", err)
	}
}
