package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ResponseError represent the reseponse error struct

func (e *ResponseError) WriteToResponse(w http.ResponseWriter) {
	w.WriteHeader(e.HTTPStatus)
	fmt.Fprint(w, e.ToJSON())
}

func (e *ResponseError) ToJSON() string {
	j, err := json.Marshal(e)
	if err != nil {
		return `{"code":50001,"message":"unable to marshal error"}`
	}
	return string(j)
}

type ResponseMessage struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseError struct {
	HTTPStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}
