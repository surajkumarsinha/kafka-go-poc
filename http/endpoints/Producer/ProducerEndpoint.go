package producer

import (
	"net/http"
)

type ProducerEndpoint struct{}

func (r *ProducerEndpoint) Produce(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("\"live\": \"ok\""))
}
