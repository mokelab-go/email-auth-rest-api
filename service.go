package emailAuthRestAPI

import (
	"encoding/json"
	"net/http"
)

// Service is an interface for REST API
type Service interface {
	Request(email string, params map[string]interface{}) Response
	Token(req, code string) Response
	Refresh(token string) Response
}

// Response describes HTTP response
type Response struct {
	Status  int
	Headers map[string]string
	Body    map[string]interface{}
}

func (r *Response) Write(w http.ResponseWriter) {
	if r.Status == http.StatusNoContent {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if b, err := json.Marshal(r.Body); err == nil {
		if r.Headers != nil {
			h := w.Header()
			for key, value := range r.Headers {
				h.Set(key, value)
			}
		}
		w.WriteHeader(r.Status)
		w.Write(b)
	}
}
