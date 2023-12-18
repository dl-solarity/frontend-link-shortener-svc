package handlers

import (
	"net/http"
	"time"

	"github.com/dl-solarity/frontend-link-shortener-svc/internal/data"
	"github.com/dl-solarity/frontend-link-shortener-svc/internal/service/requests"
	"github.com/dl-solarity/frontend-link-shortener-svc/resources"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	linkLength = 8
	padding    = 2
)

func CreateLink(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateLinkRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	path, value := request.Data.Attributes.Path, request.Data.Attributes.Value
	linkHash := getHash(getHash(path) + getHash(string(value)))

	linkData := data.Link{
		ID:        linkHash[padding : padding+linkLength],
		CreatedAt: time.Now().UTC(),
		Value:     request.Data.Attributes.Value,
		Path:      path,
	}

	var link *data.Link
	err = LinksQ(r).Transaction(func(q data.LinksQ) error {
		if err := q.Delete(linkData.ID); err != nil {
			return errors.Wrap(err, "failed to delete a link")
		}

		link, err = q.Insert(linkData)
		if err != nil {
			return errors.Wrap(err, "failed to insert a link")
		}

		return nil
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to create a link")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := resources.ShortLinkResponse{
		Data: newLinkModel(link),
	}

	ape.Render(w, response)
}

func newLinkModel(link *data.Link) resources.ShortLink {
	return resources.ShortLink{
		Key: resources.Key{
			ID:   link.ID,
			Type: resources.LINK,
		},
		Attributes: resources.ShortLinkAttributes{
			CreatedAt: link.CreatedAt,
			Value:     link.Value,
			Path:      link.Path,
		},
	}
}

func getHash(s string) string {
	hash := crypto.Keccak256Hash([]byte(s))
	return hash.Hex()
}
