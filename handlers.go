package emailAuthRestAPI

import (
	"net/http"

	"github.com/mokelab-go/hop"
)

func request(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := hop.BodyJSON(r.Context())
		email, _ := params["email"].(string)
		resp := service.Request(email, params)
		resp.Write(w)
	}
}

func token(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := hop.BodyJSON(r.Context())
		reqCode, _ := params["request"].(string)
		code, _ := params["code"].(string)
		resp := service.Token(reqCode, code)
		resp.Write(w)
	}
}

func refresh(service Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := hop.BodyJSON(r.Context())
		token, _ := params["token"].(string)
		resp := service.Refresh(token)
		resp.Write(w)
	}
}
