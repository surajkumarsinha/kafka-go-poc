package providers

import (
	generalEndpoints "github.com/surajkumarsinha/kafka-go-poc/http/endpoints"
	producerEndpoint "github.com/surajkumarsinha/kafka-go-poc/http/endpoints/producer"
	infra "github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	generalEndpointInterfaces "github.com/surajkumarsinha/kafka-go-poc/types/interfaces/endpoints"
	producerEndpointInterfaces "github.com/surajkumarsinha/kafka-go-poc/types/interfaces/endpoints/producer"
)

func EndpointProviders() {
	infra.Register(
		func(params ...any) producerEndpointInterfaces.IProducerEndpoint {
			return producerEndpoint.ProducerEndpoint{}
		},
	)
	infra.Register(
		func(a ...any) generalEndpointInterfaces.IGeneralEndpoints {
			return &generalEndpoints.GeneralEndpoints{}
		},
	)
}
