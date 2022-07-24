package domain

import (
	"context"

	"game.typing-guru.com/grpc-interfaces/game.typing-guru.com/rpc/user"
)

type domainI struct {
	userClient user.UserClient
}

func (d *domainI) GetUserById(ctx context.Context, userId string) (*user.GetUserOut, error) {
	return d.userClient.GetUserById(ctx, &user.GetUserByIdIn{
		Id: userId,
	})

}

func fxRPCClient(userClient user.UserClient) Domain {
	return &domainI{
		userClient: userClient,
	}
}
