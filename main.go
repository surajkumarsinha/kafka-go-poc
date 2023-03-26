package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"github.com/surajkumarsinha/kafka-go-poc/config"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure/interfaces"
	"github.com/surajkumarsinha/kafka-go-poc/routing"
)

func main() {
	infrastructure.KernelBuilder().Build(config.App())
	routingSystem := infrastructure.Resolve[interfaces.IChiRouter]()
	http.ListenAndServe(getAddress(), routingSystem.InitRouter(routing.Routes))
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
}
