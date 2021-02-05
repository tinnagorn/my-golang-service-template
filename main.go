package main

import (
	"runtime"
	"strings"
	"time"

	"github.com/spf13/viper"
	log "github.com/tOnkowzl/libs/logx"
	"github.com/tinnagorn/my-golang-service-template/cachemanager"
	"github.com/tinnagorn/my-golang-service-template/database"
	"github.com/tinnagorn/my-golang-service-template/health"
	"github.com/tinnagorn/my-golang-service-template/inquirydata"
	"github.com/tinnagorn/my-golang-service-template/router"
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
	err := cachemanager.NewRedisClient()
	if err != nil {
		log.Fatalf("Can't initial to Redis : %s\n", err.Error())
	}

	var routers = router.NewEcho()

	healthService := health.NewService()
	healthHandler := health.NewHandler(healthService)
	routers.GET("/health", healthHandler.HealthCheck)

	inqDataService := inquirydata.NewService()
	inqDataHandler := inquirydata.NewHandler(inqDataService)
	routers.POST("/inquiry-data", inqDataHandler.InquiryData)

	go router.Runs(routers)

	router.Shutdown(routers)
}
