package grcpserver

import (
	"context"
	er "errors"
	"testing"

	"github.com/Craftbec/Shortener_link/internal/errors"
	pb "github.com/Craftbec/Shortener_link/internal/linkshorter"
	"github.com/Craftbec/Shortener_link/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestServerAPI_Get(t *testing.T) {
	mock := storage.NewInMemory()
	server := ServerAPI{st: mock}
	ctx := context.TODO()
	t.Run("Checking the length of a short link", func(t *testing.T) {
		short := &pb.ShortLink{Link: "short"}
		ers := errors.IncorrectLength
		link, err := server.Get(ctx, short)
		assert.Error(t, err)
		assert.Nil(t, link)
		assert.True(t, er.Is(err, ers))
	})

	t.Run("should return error for invalid characters", func(t *testing.T) {
		invalid := &pb.ShortLink{Link: "invalid12+"}
		ers := errors.InvalidCharacters
		link, err := server.Get(ctx, invalid)
		assert.Error(t, err)
		assert.Nil(t, link)
		assert.True(t, er.Is(err, ers))
	})
}

func TestServerAPI_Post(t *testing.T) {
	mock := storage.NewInMemory()
	server := ServerAPI{st: mock}
	ctx := context.TODO()
	t.Run("should return error for empty link", func(t *testing.T) {
		link := &pb.OriginalLink{Link: ""}
		ers := errors.NoURL
		shortLink, err := server.Post(ctx, link)
		assert.Error(t, err)
		assert.Nil(t, shortLink)
		assert.True(t, er.Is(err, ers))
	})

	t.Run("Gives away the old short link", func(t *testing.T) {
		originalLink := &pb.OriginalLink{Link: "https://ozon.ru"}
		shortLink, err := server.Post(ctx, originalLink)
		assert.NoError(t, err)
		resultLink, er := mock.CheckPost(ctx, "https://ozon.ru")
		assert.NoError(t, er)
		assert.Equal(t, shortLink.Link, resultLink)
		shortLink2, _ := server.Post(ctx, originalLink)
		assert.Equal(t, shortLink.Link, shortLink2.Link)
	})

}
