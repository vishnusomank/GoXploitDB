package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vishnusomank/GoXploitDB/fetcher"
	"github.com/vishnusomank/GoXploitDB/models"
	"github.com/vishnusomank/GoXploitDB/server"
)

func main() {

	// Getting application configuration details from conf/ path
	configFilePath := flag.String("config-path", "conf/", "conf/")
	flag.Parse()

	//Load application configuration
	server.LoadConfig(configFilePath)

	go fetcher.StartGit()

	//Connect to SqLite DB
	models.ConnectDatabase()
	r := gin.New()
	//Allow CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	server.SetupRoutes(r)
	fmt.Printf("[%s][%s] Configurations loaded\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))
	fmt.Printf("[%s][%s] Database configured\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))
	fmt.Printf("[%s][%s] Routes loaded\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))
	fmt.Printf("[%s][%s] Server started\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))
	server.StartServer(r)

}
