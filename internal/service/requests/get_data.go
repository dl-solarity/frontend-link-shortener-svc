package requests

import (
	"net/http"

	"github.com/go-chi/chi"
)

type GetDataRequest struct {
	Link string `url:"-"`
}

func NewGetDataRequest(r *http.Request) (GetDataRequest, error) {
	var request GetDataRequest
	request.Link = chi.URLParam(r, "link")

	return request, nil
}
