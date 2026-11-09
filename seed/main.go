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

func defineTheme(win *ui.Window, color uint32) {
	win.DefineTheme(ui.NewThemeFromSeed(color, false))
	win.DefineTheme(ui.NewThemeFromSeed(color, true))
}

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	win, err := app.LoadWindowFile(
		"seed/seed.goi",
		ui.WithHotReload(true),
		ui.WithSize(520, 620),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	seed := uint32(0x411F89FF)

	defineTheme(win, seed)

	if err := win.Bind(func(w *ui.Window) error {
		ui.SetHandler(w, "onSetSeedColor", func(color uint32) {
			seed = color
			defineTheme(w, seed)
		})

		ui.SetHandler(w, "onToggleScheme", func(isDark bool) {
			if isDark {
				w.Root().Theme.Set("darkFromSeed")
				return
			}
			w.Root().Theme.Set("lightFromSeed")
		})

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
