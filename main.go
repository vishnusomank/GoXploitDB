package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vishnusomank/GoXploitDB/constants"
	"github.com/vishnusomank/GoXploitDB/services"
)

var configFilePath *string

func main() {

	// Getting application configuration details from conf/ path
	configFilePath = flag.String("config-path", "conf/", "conf/")
	flag.Parse()

	//Load application configuration
	loadConfig()

	CURRENT_DIR, err := os.Getwd()
	if err != nil {
		fmt.Printf("[%s][%s] Failed to get current directory: %v\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"), err)

	}

	GIT_DIR := CURRENT_DIR + "/exploitdb"

	services.Git_Operation(GIT_DIR)

	//Connect to SqLite DB
	//models.ConnectDatabase()
	r := gin.New()
	//Allow CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	setupRoutes(r)
	fmt.Printf("[%s][%s] Configurations loaded\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))
	fmt.Printf("[%s][%s] Database configured\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))
	fmt.Printf("[%s][%s] Routes loaded\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.GreenString("DONE"))
	startServer(r)

}
func loadConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		if readErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("[%s][%s] No config file found at %s\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"), *configFilePath)
		} else {
			fmt.Printf("[%s][%s] Error reading config file: %s\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"), readErr)
		}
	}
}

// startServer - Start server
func startServer(r *gin.Engine) {
	srv := &http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("[%s][%s] Listen: %s\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"), err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("[%s][%s] Shutting down server...\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"))

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("[%s][%s] Server forced to shutdown: %s\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"), err)

	}

	fmt.Printf("[%s][%s] Server exiting\n", color.BlueString(time.Now().Format("01-02-2006 15:04:05")), color.RedString("ERR"))

}

//setupRoutes -
func setupRoutes(r *gin.Engine) {

	application := r.Group(viper.GetString("server.basepath"))
	{
		v1 := application.Group("/api/v1")
		{
			v1.POST(constants.CVE, services.SearchByCVE)

		}
	}
}
