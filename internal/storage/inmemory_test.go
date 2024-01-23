package storage

import (
	"context"
	"testing"

	"github.com/Craftbec/Shortener_link/internal/errors"
	"github.com/Craftbec/Shortener_link/internal/shorting"
)

func TestInMemoryGet(t *testing.T) {
	bd := NewInMemory()
	originalLink := "https://ozon.ru"
	shortLink := shorting.GenerateShortLink()
	bd.shortLink[shortLink] = originalLink
	bd.originalLink[originalLink] = shortLink
	resultOriginal, err := bd.Get(context.Background(), shortLink)
	if err != nil {
		t.Errorf("Error getting original link")
	}
	if resultOriginal != originalLink {
		t.Errorf("Incorrect original link received")
	}
}

func TestInMemoryPostAndCheckPost(t *testing.T) {
	bd := NewInMemory()
	originalLink := "https://ozon.ru"
	shortLink := shorting.GenerateShortLink()
	err := bd.Post(context.Background(), originalLink, shortLink)
	if err != nil {
		t.Errorf("The link was not added")
	}
	resultShortLink, err := bd.CheckPost(context.Background(), originalLink)
	if err != nil {
		t.Errorf("Error when searching by link")
	}
	noLink := "https://test.com"
	_, err = bd.CheckPost(context.Background(), noLink)
	if err != errors.NotFound {
		t.Errorf("Error searching for non-existent link")
	}
	if resultShortLink != shortLink {
		t.Errorf("The short link is incorrect in relation to the original")
	}
}
