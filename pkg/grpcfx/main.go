package grpcfx

import (
	"context"
	"fmt"
	"net"
	"os"

	"game.typing-guru.com/grpc-interfaces/game.typing-guru.com/rpc/user"
	"game.typing-guru.com/pkg/errors"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewInsecureClient(grpcUrl string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(grpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

var FxGrpcClient = fx.Module(
	"grpc-client",
	fx.Provide(
		func() (*grpc.ClientConn, error) {
			grpcUrl := os.Getenv("GRPC_URL")
			if grpcUrl == "" {
				panic("GRPC_URL environment not provided")
			}
			return NewInsecureClient(grpcUrl)
		},
	),
	fx.Invoke(
		func(grpcClient *grpc.ClientConn, lifecycle fx.Lifecycle) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return nil
					},
				},
			)
		},
	),
	fx.Provide(
		func(conn *grpc.ClientConn) user.UserClient {
			return user.NewUserClient((*grpc.ClientConn)(conn))
		},
	),
)

var FxGrpcServer = fx.Module("grpc-server",
	fx.Provide(grpc.NewServer),
	fx.Invoke(func(lf fx.Lifecycle, server *grpc.Server) {
		GRPCPort := os.Getenv("GRPC_PORT")
		if GRPCPort == "" {
			panic("GRPC_PORT environment not provided")
		}
		lf.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				listen, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPCPort))
				if err != nil {
					return errors.NewEf(err, "could not listen to net/tcp server")
				}
				go func() error {
					err := server.Serve(listen)
					if err != nil {
						return errors.NewEf(err, "could not start grpc server ")
					}
					return nil
				}()
				return nil
			},
			OnStop: func(context.Context) error {

				fmt.Print("stopped grpc")
				server.Stop()
				return nil
			},
		})
	}),
)
