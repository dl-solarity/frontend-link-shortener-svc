package handlers

import (
	"context"
	"net/http"

	"github.com/dl-solarity/frontend-link-shortener-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	linksQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxLinksQ(entry data.LinksQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, linksQCtxKey, entry)
	}
}

func LinksQ(r *http.Request) data.LinksQ {
	return r.Context().Value(linksQCtxKey).(data.LinksQ).New()
}
