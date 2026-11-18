// Seed: one colour, a whole scheme. The app builds a light and a dark theme
// from the same seed with ui.NewThemeFromSeed and defines both by name, so the
// design selects one with `theme: lightFromSeed` / `theme: darkFromSeed`.
//
// Run with: make seed (or cd seed && goi run)
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	win, err := app.LoadWindow(
		"Main.goi",
		ui.WithSize(520, 620),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	seed := uint32(0x411F89FF)

	defineTheme(win, seed)

	hidePicker := func(w *ui.Window) {
		if picker, err := w.ByID[*ui.ColorPicker]("seedPicker"); err == nil {
			picker.SetVisible(false)
		}
	}

	if err := win.Bind(func(w *ui.Window) error {
		ui.SetHandler(w, "onSetSeedColor", func(hex string) {
			color, ok := parseSeed(hex)
			if !ok {
				return
			}
			seed = color
			defineTheme(w, seed)
			hidePicker(w)
		})

		ui.SetHandler(w, "onOpenSeedPicker", func() {
			picker, err := w.ByID[*ui.ColorPicker]("seedPicker")
			if err != nil {
				return
			}
			picker.Text.Set(fmt.Sprintf("#%06x", seed>>8))
			picker.SetVisible(true)
		})

		ui.SetHandler(w, "onCloseSeedPicker", func() { hidePicker(w) })

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

// parseSeed reads the #rrggbb or #rrggbbaa string a ColorPicker reports, and
// the shorthand preset values as #rrggbb.
func parseSeed(hex string) (uint32, bool) {
	h := strings.TrimPrefix(hex, "#")
	if len(h) == 6 {
		h += "ff"
	}
	if len(h) != 8 {
		return 0, false
	}
	v, err := strconv.ParseUint(h, 16, 32)
	if err != nil {
		return 0, false
	}
	return uint32(v), true
}
