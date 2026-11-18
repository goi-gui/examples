// Command counter is the framework's smoke test: a window loaded from a .goi
// file, buttons whose handlers update a label through the typed property API.
//
// Run with: make counter (or cd counter && goi run)
package main

import (
	"github.com/goi-gui/goi/ui"
)

func main() {
	app, err := ui.NewApp()
	if err != nil {
		panic(err)
	}

	win, err := app.LoadWindow("ui/Main.goi")
	if err != nil {
		panic(err)
	}

	// store count variable outside of the binding,
	// if you don't want reset after hot reload
	count := 0

	err = win.Bind(func(w *ui.Window) error {
		ui.SetData(w, "Count", count)
		ui.SetHandler(w, "onInc", func() {
			count++
			ui.SetData(w, "Count", count)
		})
		ui.SetHandler(w, "onDec", func() {
			count--
			ui.SetData(w, "Count", count)
		})
		ui.SetHandler(w, "onReset", func() {
			count = 0
			ui.SetData(w, "Count", count)
		})
		return nil
	})
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
