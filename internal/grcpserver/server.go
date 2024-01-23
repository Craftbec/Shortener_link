package grcpserver

import (
	"context"
	er "errors"
	"fmt"
	"net"
	"strings"

	"github.com/Craftbec/Shortener_link/config"
	"github.com/Craftbec/Shortener_link/internal/errors"
	pb "github.com/Craftbec/Shortener_link/internal/linkshorter"
	"github.com/Craftbec/Shortener_link/internal/shorting"
	"github.com/Craftbec/Shortener_link/internal/storage"
	"google.golang.org/grpc"
)

type ServerAPI struct {
	pb.UnimplementedShortenerServer
	st storage.Storage
}

func (s *ServerAPI) Get(ctx context.Context, link *pb.ShortLink) (*pb.OriginalLink, error) {
	shortLink := link.Link
	if len(shortLink) != shorting.Size {
		return nil, errors.IncorrectLength
	}
	if len(strings.Trim(shortLink, shorting.Alphabet)) != 0 {
		return nil, errors.InvalidCharacters
	}
	originalLink, err := s.st.Get(ctx, shortLink)
	if err != nil {
		if er.Is(err, errors.NotFound) {
			return nil, err
		}
		return nil, errors.InternalServerError
	}
	return &pb.OriginalLink{Link: originalLink}, nil
}

func (s *ServerAPI) Post(ctx context.Context, link *pb.OriginalLink) (*pb.ShortLink, error) {
	if len(link.Link) == 0 {
		return nil, errors.NoURL
	}
	shortLink, err := s.st.CheckPost(ctx, link.Link)
	if err == nil {
		return &pb.ShortLink{Link: shortLink}, nil
	}
	shortLink = shorting.GenerateShortLink()
	_, err = s.st.Get(ctx, shortLink)
	for !er.Is(err, errors.NotFound) {
		shortLink = shorting.GenerateShortLink()
		_, err = s.st.Get(ctx, shortLink)
	}
	if err != nil && !er.Is(err, errors.NotFound) {
		return nil, errors.InternalServerError
	}
	err = s.st.Post(ctx, link.Link, shortLink)
	if err != nil {
		return nil, errors.InternalServerError
	}
	return &pb.ShortLink{Link: shortLink}, nil
}

func GRPCServer(ctx context.Context, st storage.Storage, conf *config.Config) error {
	s := grpc.NewServer()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GRCP.Port))
	if err != nil {
		return err
	}
	server := ServerAPI{st: st}
	pb.RegisterShortenerServer(s, &server)
	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()
	if err := s.Serve(l); err != nil {
		return err
	}
	return nil
}
