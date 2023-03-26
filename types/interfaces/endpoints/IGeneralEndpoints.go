package endpoints

import "net/http"

type IGeneralEndpoints interface {
	Health(res http.ResponseWriter, req *http.Request)
}
