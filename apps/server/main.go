package main

import (
	"game.typing-guru.com/apps/server/internal/framework"
	"go.uber.org/fx"
)

func main() {
	fx.New(framework.Module).Run()
}
