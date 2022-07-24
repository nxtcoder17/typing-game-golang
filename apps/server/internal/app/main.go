package app

import (
	"game.typing-guru.com/apps/server/internal/domain"
	"go.uber.org/fx"
)

var Module = fx.Module("app", domain.Module)
