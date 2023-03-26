package producer

import "net/http"

type IProducerEndpoint interface {
	Produce(res http.ResponseWriter, req *http.Request)
}
