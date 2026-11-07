// Translations: t(...) reads the catalogs in i18n/ and the buttons switch the
// language from Go. Editing a JSON file under i18n/ updates the text live.
//
// Run with: go run -ldflags=-linkmode=internal ./i18n
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
		"i18n/Main.goi",
		ui.WithHotReload(true),
		ui.WithSize(520, 460),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Both defaults follow the desktop; English is the fallback for the keys a
	// language leaves out. This is app state, so a reload does not reset it.
	app.SetFallbackLanguage("en")

	if err := win.Bind(func(w *ui.Window) error {
		languages := []struct{ id, code string }{
			{"btnEn", "en"},
			{"btnTr", "tr"},
			{"btnDe", "de"},
		}
		for _, l := range languages {
			l := l
			if b, err := ui.ByID[*ui.Button](w, l.id); err == nil {
				b.OnClick(func(ui.ClickEvent) { app.SetLanguage(l.code) })
			}
		}
		// An empty code follows the desktop again.
		if b, err := ui.ByID[*ui.Button](w, "btnSystem"); err == nil {
			b.OnClick(func(ui.ClickEvent) { app.SetLanguage("") })
		}
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
