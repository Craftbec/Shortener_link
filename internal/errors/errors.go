package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	NotFound            = status.Error(codes.NotFound, "Not found")
	IncorrectLength     = status.Error(codes.InvalidArgument, "Incorrect short link length")
	InvalidCharacters   = status.Error(codes.InvalidArgument, "Short link contains invalid characters")
	InternalServerError = status.Error(codes.Internal, "Internal Server Error")
	NoURL               = status.Error(codes.InvalidArgument, "No URL entered")
)
