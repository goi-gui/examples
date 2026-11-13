// Seed: one colour, a whole scheme. The app builds a light and a dark theme
// from the same seed with ui.NewThemeFromSeed and defines both by name, so the
// design selects one with `theme: lightFromSeed` / `theme: darkFromSeed`.
//
// Run with: go run -ldflags=-linkmode=internal ./seed
package main

import (
	"fmt"
	"os"

	"github.com/goi-gui/goi/ui"
)

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	win, err := app.LoadWindowFile(
		"markdown/markdown.goi",
		ui.WithHotReload(true),
		ui.WithSize(520, 620),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := win.Bind(func(w *ui.Window) error {
		return nil
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
