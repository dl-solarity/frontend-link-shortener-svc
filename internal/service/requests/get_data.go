package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/urlval"
)

type GetDataRequest struct {
	Link string `url:"-"`
}

func NewGetDataRequest(r *http.Request) (GetDataRequest, error) {
	var request GetDataRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	request.Link = chi.URLParam(r, "link")

	return request, nil
}
