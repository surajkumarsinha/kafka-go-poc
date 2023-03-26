package producer

import (
	"github.com/go-chi/chi/v5"
	"github.com/surajkumarsinha/kafka-go-poc/infrastructure"
	producerEndpoint "github.com/surajkumarsinha/kafka-go-poc/types/interfaces/endpoints/producer"
)

type ProducerEndpoint struct{}

func (ep *ProducerEndpoint) Routes(r *chi.Mux) {
	producerEndpoint := infrastructure.Resolve[producerEndpoint.IProducerEndpoint]()

	r.Post("/producer/produce", producerEndpoint.Produce)
}
