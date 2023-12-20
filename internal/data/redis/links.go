package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dl-solarity/frontend-link-shortener-svc/internal/data"
	"github.com/redis/go-redis/v9"
)

type linksQ struct {
	client *redis.Client
}

func NewLinksQ(client *redis.Client) data.LinksQ {
	return &linksQ{client: client}
}

func (q *linksQ) New() data.LinksQ {
	return NewLinksQ(q.client)
}

func (q *linksQ) Get(ctx context.Context, id string) (*data.Link, error) {
	val, err := q.client.Get(ctx, id).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var link data.Link
	err = json.Unmarshal([]byte(val), &link)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (q *linksQ) Insert(ctx context.Context, value data.Link) (*data.Link, error) {
	val, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	err = q.client.Set(ctx, value.ID, val, value.ExpiredAt.Sub(time.Now())).Err()
	if err != nil {
		return nil, err
	}

	return &value, nil
}
