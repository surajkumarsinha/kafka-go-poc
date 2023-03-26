package endpoints

import "net/http"

type GeneralEndpoints struct{}

func (ge *GeneralEndpoints) Health(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("ok"))
}
