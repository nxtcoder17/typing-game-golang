package domain

import (
	"game.typing-guru.com/grpc-interfaces/game.typing-guru.com/rpc/user"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Module("domain",
	fx.Provide(fxRPCServer),
	fx.Invoke(func(server *grpc.Server, userServer user.UserServer) {
		user.RegisterUserServer(server, userServer)
	}),
)
