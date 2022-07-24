package domain

import (
	"context"
	"game.typing-guru.com/grpc-interfaces/game.typing-guru.com/rpc/user"
)

type Domain interface {
	GetUserById(ctx context.Context, userId string) (*user.GetUserOut, error)
}
