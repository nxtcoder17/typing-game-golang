package ui

import (
	"context"
	"fmt"

	"game.typing-guru.com/apps/client/internal/domain"
	"go.uber.org/fx"
)

var Module = fx.Module("ui",
	fx.Invoke(func(d domain.Domain) {
		ctx := context.TODO()
		resp, err := d.GetUserById(ctx, "hi")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(resp)
	}),
	HomeModule,
)
