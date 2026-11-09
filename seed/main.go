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

var seeds = []uint32{0x6750A4FF, 0x2E7D32FF, 0xE8710AFF}

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

	// State lives outside Bind, because a hot reload runs Bind again.
	dark := false
	seed := seeds[0]

	// Both moods are built from the same seed. Their names come from the scheme
	// itself, and redefining a name replaces it, which is how a new seed lands.
	define := func() {
		win.DefineTheme(ui.NewThemeFromSeed(seed, false))
		win.DefineTheme(ui.NewThemeFromSeed(seed, true))
	}
	define()

	if err := win.Bind(func(w *ui.Window) error {
		apply := func() {
			name := "lightFromSeed"
			if dark {
				name = "darkFromSeed"
			}
			w.Root().Theme.Set(name)
			if out, err := ui.ByID[*ui.Text](w, "current"); err == nil {
				out.Text.Set(name)
			}
		}
		for i, id := range []string{"seedPurple", "seedGreen", "seedOrange"} {
			i := i
			if b, err := ui.ByID[*ui.Button](w, id); err == nil {
				b.OnClick(func(ui.ClickEvent) {
					seed = seeds[i]
					define()
					apply()
				})
			}
		}
		if sw, err := ui.ByID[*ui.Switch](w, "darkSwitch"); err == nil {
			sw.OnToggle(func(on bool) {
				dark = on
				apply()
			})
		}
		apply()
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
