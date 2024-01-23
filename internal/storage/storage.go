package storage

import "context"

type Storage interface {
	Get(ctx context.Context, short string) (string, error)
	Post(ctx context.Context, original string, short string) error
	CheckPost(ctx context.Context, original string) (string, error)
	GracefulStopDB()
}
