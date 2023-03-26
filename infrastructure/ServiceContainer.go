package infrastructure

import (
	"fmt"
	"reflect"
	"sync"
)

type service[T any] struct {
	impl func(...any) T
}

var services map[string]any

func Register[T any](implement func(...any) T) {
	typeName := reflect.TypeOf(new(T)).String()

	if _, ok := services[typeName]; ok {
		panic(fmt.Sprintf("service [%s] already registered", typeName))
	}

	services[typeName] = service[T]{impl: implement}
}

func Resolve[T any](params ...any) T {
	if !isServiceContainerInitialised {
		panic("service container is not initialised")
	}

	typeName := reflect.TypeOf(new(T)).String()
	serv, ok := services[typeName]

	if ok {
		return serv.(service[T]).impl(params...)
	}
	panic(fmt.Sprintf("unregistered service [%s]", typeName))
}

var (
	isServiceContainerInitialised bool
	serviceContainerOnce          sync.Once
)

func InitializeServiceContainer(config map[string]any) {
	if !isServiceContainerInitialised {
		serviceContainerOnce.Do(func() {
			services = map[string]any{}

			providers := config["providers"].([]func())

			for _, provider := range providers {
				provider()
			}

			isServiceContainerInitialised = true
		})
	}
}
