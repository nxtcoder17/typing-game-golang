package app

import (
	"game.typing-guru.com/apps/client/internal/domain"
	"game.typing-guru.com/apps/client/internal/ui"
	"go.uber.org/fx"
)

var Module = fx.Module("app",

	domain.Module,
	ui.Module,
)
