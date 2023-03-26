package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"github.com/surajkumarsinha/kafka-go-poc/config"
	"github.com/surajkumarsinha/kafka-go-poc/http/routing"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure/interfaces"
)

func main() {
	infrastructure.KernelBuilder().Build(config.App())
	http.ListenAndServe(getAddress(), getRoutingHandler())
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
}

func getRoutingHandler() http.Handler {
	routingSystem := infrastructure.Resolve[interfaces.IChiRouter]()
	routes := routing.InitRoutes().Routes
	return routingSystem.InitRouter(routes)
}
