package framework

import (
	"game.typing-guru.com/apps/server/internal/app"
	"game.typing-guru.com/pkg/grpcfx"
	"go.uber.org/fx"
)

var Module = fx.Module("framework",
	grpcfx.FxGrpcServer,
	app.Module,
)
