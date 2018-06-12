package emailAuthRestAPI

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mokelab-go/hop"
)

// InitRouter is an entry point
func InitRouter(router *mux.Router, service Service) {
	router.Methods(http.MethodPost).
		Path("/auth/request").
		Handler(hop.Operations(
			hop.GetBodyAsJSON,
		)(request(service)))
	router.Methods(http.MethodPost).
		Path("/auth/token").
		Handler(hop.Operations(
			hop.GetBodyAsJSON,
		)(token(service)))
	router.Methods(http.MethodPost).
		Path("/auth/refresh").
		Handler(hop.Operations(
			hop.GetBodyAsJSON,
		)(refresh(service)))
}
