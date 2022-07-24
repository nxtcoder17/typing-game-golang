package ui

import (
	"context"
	"fmt"

	"game.typing-guru.com/apps/client/internal/domain"
	"github.com/rivo/tview"
	"go.uber.org/fx"
)

var HomeModule = fx.Module("home",
	fx.Invoke(func(d domain.Domain) {
		ctx := context.TODO()
		resp, err := d.GetUserById(ctx, "hi")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(resp)

		app := tview.NewApplication()

		textview := tview.NewTextView().SetText(resp.Email)

		if err := app.SetRoot(textview, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}

	}),
)
