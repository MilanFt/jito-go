package relayer_client

import (
	"github.com/MilanFt/jito-go/pkg"
	"github.com/MilanFt/jito-go/proto"
	"google.golang.org/grpc"
)

type Client struct {
	GrpcConn *grpc.ClientConn

	Relayer proto.RelayerClient

	Auth *pkg.AuthenticationService

	ErrChan <-chan error // ErrChan is used for dispatching errors from functions executed within goroutines.
}
