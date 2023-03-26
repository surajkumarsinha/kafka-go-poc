package providers

import (
	producerEndpoint "github.com/surajkumarsinha/kafka-go-poc/http/endpoints/producer"
	infra "github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	producerEndpointInterfaces "github.com/surajkumarsinha/kafka-go-poc/types/interfaces/endpoints/producer"
)

func EndpointProviders() {
	infra.Register(
		func(params ...any) producerEndpointInterfaces.IProducerEndpoint {
			return producerEndpoint.ProducerEndpoint{}
		},
	)
}
