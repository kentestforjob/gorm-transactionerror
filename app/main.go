package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"test/gormtransactionerr/app/di"
	"test/gormtransactionerr/app/repositories/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var route_engine *gin.Engine

// var dbConn *gorm.DB

type IndexData struct {
	Title   string
	Content string
}

func init() {

	myOS, myArch := runtime.GOOS, runtime.GOARCH
	inContainer := "inside"
	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		inContainer = "outside"
	}
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.SetConfigName(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("APP_DEBUG") {
		gin.SetMode(gin.DebugMode)
		log.Println("Service RUN on DEBUG mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// database config or other config will store in app.json
	viper.SetConfigType("json")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")

	target_config_file := "app"
	if viper.Get("APP_ENV") == "local" {
		target_config_file = "local_" + target_config_file
	}

	viper.SetConfigName(target_config_file)
	err = viper.MergeInConfig()

	if err != nil {
		panic(err)
	}

	fmt.Printf("I'm running on %s/%s.\n", myOS, myArch)
	fmt.Printf("I'm running %s of a container.\n", inContainer)
	// confirm where the file has been read in from
	fmt.Println("ConfigFileUsed ", viper.ConfigFileUsed())

}

func main() {

	// Set the router as the default one provided by Gin
	route_engine = gin.Default()

	// // only for openAPI doc. Need to comment it after developement
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	// corsConfig.AllowHeaders = []string{"*"}
	// //corsConfig.AllowOrigins = []string{"*"}
	// corsConfig.AllowAllOrigins = true
	// corsConfig.AllowCredentials = true
	// route_engine.Use(cors.New(corsConfig))

	// initial database connection - mysql
	dbConn := db.ConnectMysqlGormDatabase()

	// depdency injection
	di.InitializeAPIs(dbConn, route_engine)

	// Start serving the application
	route_engine.Run(viper.GetString("server.address"))
}
