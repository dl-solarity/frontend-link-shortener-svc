package requests

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/dl-solarity/frontend-link-shortener-svc/resources"
	. "github.com/go-ozzo/ozzo-validation"
)

type CreateLinkRequest struct {
	Data resources.CreateShortLink
}

func NewCreateLinkRequest(r *http.Request) (CreateLinkRequest, error) {
	var request CreateLinkRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, validateCreateLinkRequest(request)
}

func validateCreateLinkRequest(request CreateLinkRequest) error {
	attrs := &request.Data.Attributes

	return ValidateStruct(attrs,
		Field(&attrs.Value, Required),
		Field(&attrs.Path, Required),
	)
}
