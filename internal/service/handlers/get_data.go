package handlers

import (
	"net/http"

	"github.com/dl-solarity/frontend-link-shortener-svc/internal/service/requests"
	"github.com/dl-solarity/frontend-link-shortener-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetDataRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	link, err := LinksQ(r).Get(r.Context(), request.Link)
	if err != nil {
		Log(r).WithError(err).Error("failed to get a link")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if link == nil {
		Log(r).Warn("link not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	response := resources.ShortLinkResponse{
		Data: newLinkModel(link),
	}

	ape.Render(w, response)
}
