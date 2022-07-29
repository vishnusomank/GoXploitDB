package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vishnusomank/GoXploitDB/services"
	"github.com/vishnusomank/GoXploitDB/utils"
)

func LoadConfig(configFilePath *string) {
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
func StartServer(r *gin.Engine) {
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
func SetupRoutes(r *gin.Engine) {

	application := r.Group(viper.GetString("server.basepath"))
	{
		v1 := application.Group("/api/v1")
		{
			v1.GET(utils.CVE, services.SearchByCVE)
			v1.GET(utils.PLATFORM, services.SearchByPlatform)
			v1.GET(utils.TYPE, services.SearchByType)

			v1.GET(utils.All_PLATFORM, services.ShowAllPlatform)
			v1.GET(utils.ALL_TYPE, services.ShowAllType)

		}
	}
}
