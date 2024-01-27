package storage

import (
	"context"
	"sync"

	"github.com/Craftbec/Shortener_link/internal/errors"
)

type Storage interface {
	Get(ctx context.Context, short string) (string, error)
	Post(ctx context.Context, original string, short string) error
	CheckPost(ctx context.Context, original string) (string, error)
	GracefulStopDB()
}

type InMemory struct {
	mu           sync.RWMutex
	originalLink map[string]string
	shortLink    map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{
		originalLink: make(map[string]string),
		shortLink:    make(map[string]string),
	}
}

func (i *InMemory) Get(ctx context.Context, short string) (string, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	val, ok := i.shortLink[short]
	if !ok {
		return "", errors.NotFound
	}
	return val, nil
}

func (i *InMemory) Post(ctx context.Context, original string, short string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.shortLink[short] = original
	i.originalLink[original] = short
	return nil
}

func (i *InMemory) CheckPost(ctx context.Context, original string) (string, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	val, ok := i.originalLink[original]
	if !ok {
		return "", errors.NotFound
	}
	return val, nil
}

func (i *InMemory) GracefulStopDB() {
}
