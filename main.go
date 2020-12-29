package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	log "github.com/tOnkowzl/libs/logx"
	"github.com/tOnkowzl/libs/middleware"
	"github.com/tinnagorn/my-golang-service-template/cachemanager"
	"github.com/tinnagorn/my-golang-service-template/database"
	"github.com/tinnagorn/my-golang-service-template/health"
	"github.com/tinnagorn/my-golang-service-template/inquirydata"
	"github.com/tinnagorn/my-golang-service-template/utility"
)

var (
	buildstamp = time.Now().String()
	githash    = "developing"
)
var redisClient *cachemanager.Cache

func init() {
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.MergeConfig(strings.NewReader(viper.GetString("configs"))); err != nil {
		log.Panic(err.Error())
	}
}

func init() {
	runtime.GOMAXPROCS(1)
	log.Init(viper.GetString("log.level"), viper.GetString("log.env"))
}

func main() {

	utility.GetSecretValue()
	database.Initialize()
	err := newRedisClient()
	if err != nil {
		log.Fatalf("Can't initial to Redis : %s\n", err.Error())
	}

	var router = newEcho()

	healthService := health.NewService()
	healthHandler := health.NewHandler(healthService)
	router.GET("/health", healthHandler.HealthCheck)

	inqDataService := inquirydata.NewService()
	inqDataHandler := inquirydata.NewHandler(inqDataService)
	router.POST("/inquiry-data", inqDataHandler.InquiryData)

	go run(router)

	shutdown(router)
}

func newEcho() *echo.Echo {
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

func run(router *echo.Echo) {
	log.Infof("starting %s", viper.GetString("app.name"))
	log.Infof("application serve at port %s", viper.GetString("app.port"))
	log.Info(router.Start(":" + viper.GetString("app.port")))
}

func shutdown(router *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	defer database.Close()
	defer redisClient.Close()
	if err := router.Shutdown(context.Background()); err != nil {
		log.Fatal("shutdown server: ", err)
	}
}

func newRedisClient() error {
	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetString("redis.port")
	redisPw := viper.GetString("secrets.redis.password")
	redisDb := viper.GetInt("redis.db")
	client, err := cachemanager.NewCache(redisHost, redisPort, redisPw, redisDb)
	if err != nil {
		log.Fatalf("Error on init Redis %s", err.Error())
		return err
	}
	redisClient = client
	return nil
}
