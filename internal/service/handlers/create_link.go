package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/dl-solarity/frontend-link-shortener-svc/internal/data"
	"github.com/dl-solarity/frontend-link-shortener-svc/internal/service/requests"
	"github.com/dl-solarity/frontend-link-shortener-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
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
		ID:        linkHash,
		CreatedAt: time.Now().UTC(),
		Value:     request.Data.Attributes.Value,
		Path:      path,
	}

	link, err := LinksQ(r).Insert(linkData)
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
	hash := sha256.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}