package domain

import (
	"context"

	"game.typing-guru.com/grpc-interfaces/game.typing-guru.com/rpc/user"
)

type domainI struct {
	user.UnimplementedUserServer
}

func (d *domainI) GetUserByEmail(_ context.Context, input *user.GetUserByEmailIn) (*user.GetUserOut, error) {

	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func (d *domainI) GetUserById(_ context.Context, input *user.GetUserByIdIn) (*user.GetUserOut, error) {

	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func (d *domainI) GetUserByName(_ context.Context, input *user.GetUserByNameIn) (*user.GetUserOut, error) {

	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func (d *domainI) GetUserBySession(_ context.Context, input *user.GetUserBySessionIdIn) (*user.GetUserOut, error) {

	return &user.GetUserOut{
		Username: "sample",
		Email:    "sample@gmail.com",
		Id:       "user_sample",
		Status:   "active",
	}, nil
}

func fxRPCServer() user.UserServer {
	return &domainI{}
}
