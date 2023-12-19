package service

import (
	"github.com/dl-solarity/frontend-link-shortener-svc/internal/data/redis"
	"github.com/dl-solarity/frontend-link-shortener-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxLinksQ(redis.NewLinksQ(s.rdb)),
			handlers.CtxLinks(*s.linksConfig),
		),
	)
	r.Route("/shortener", func(r chi.Router) {
		r.Post("/", handlers.CreateLink)
		r.Get("/{link}", handlers.GetData)
	})

	return r
}
