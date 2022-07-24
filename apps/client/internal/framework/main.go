package framework

import (
	"game.typing-guru.com/apps/client/internal/app"
	"game.typing-guru.com/pkg/grpcfx"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"framework",
	grpcfx.FxGrpcClient,
	app.Module,
)
