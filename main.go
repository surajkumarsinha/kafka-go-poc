package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"github.com/surajkumarsinha/kafka-go-poc/config"
	"github.com/surajkumarsinha/kafka-go-poc/http/routing"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure/interfaces"
)

func main() {
	infrastructure.KernelBuilder().Build(config.App())

	server := &http.Server{Addr: getAddress(), Handler: getRoutingHandler()}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, _ := context.WithTimeout(serverCtx, 5*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
}

func getRoutingHandler() http.Handler {
	routingSystem := infrastructure.Resolve[interfaces.IChiRouter]()
	routes := routing.InitRoutes().Routes
	return routingSystem.InitRouter(routes)
}
