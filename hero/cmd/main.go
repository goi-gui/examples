// Hero: one picture shared by two pages. The design holds both pages; Go only
// flips the bridge bool the `if` reads, and the Hero with the matching tag does
// the flight.
//
// Run with: make hero (or cd hero && goi run)
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

	// Outside Bind: a hot reload re-runs Bind and must not lose the page.
	detail := false

	err = win.Bind(func(w *ui.Window) error {
		ui.SetData(w, "detail", detail)

		ui.SetHandler(w, "onOpen", func() {
			detail = true
			ui.SetData(w, "detail", detail)
		})
		ui.SetHandler(w, "onClose", func() {
			detail = false
			ui.SetData(w, "detail", detail)
		})
		return nil
	})
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
