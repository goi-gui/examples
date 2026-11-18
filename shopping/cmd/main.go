// Command shopping is a product page: a rail of categories, a picture column
// and the product's details beside it.
//
// Run with: make shopping (or cd shopping && goi run)
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

	win, err := app.LoadWindow("Main.goi")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// One rose seed, both moods: the page's pinks come from the seed through the
	// primary and neutral palettes, not from per-element colours.
	const seed = 0xE91E63FF
	win.DefineTheme(ui.NewThemeFromSeed(seed, false))
	win.DefineTheme(ui.NewThemeFromSeed(seed, true))

	err = win.Bind(func(w *ui.Window) error {
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
