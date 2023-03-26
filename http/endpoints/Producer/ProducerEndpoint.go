package producer

import (
	"fmt"
	"net/http"
)

type ProducerEndpoint struct{}

func (r ProducerEndpoint) Produce(res http.ResponseWriter, req *http.Request) {
	fmt.Println("A new data is produced")
	res.Write([]byte("\"live\": \"ok\""))
}
