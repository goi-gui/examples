package main

import (
	"fmt"

	"github.com/goi-gui/goi/ui"
)

func main() {
	app, err := ui.NewApp()
	if err != nil {
		panic(err)
	}

	win, err := app.LoadWindowFile("search/search.goi", ui.WithHotReload(true))
	if err != nil {
		panic(err)
	}

	win.Bind(func(w *ui.Window) error {
		ui.SetHandler(w, "onSearch", func(term string) {
			fmt.Println("searching:", term)
		})
		return nil
	})

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
