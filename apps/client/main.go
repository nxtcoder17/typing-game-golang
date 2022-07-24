package main

import (
	"game.typing-guru.com/apps/client/internal/framework"
	"go.uber.org/fx"
)

func main() {
	fx.New(framework.Module).Run()
}
