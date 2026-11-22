// Translations: t(...) reads the catalogs in ui/i18n/ and the buttons switch the
// language from Go. Editing a JSON file under ui/i18n/ updates the text live.
//
// Run with: make i18n (or cd i18n && goi run)
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
	win, err := app.LoadWindow("ui/Main.goi")
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
			if b, err := w.ByID[*ui.Button](l.id); err == nil {
				b.OnClick(func(ui.ClickEvent) { app.SetLanguage(l.code) })
			}
		}
		// An empty code follows the desktop again.
		if b, err := w.ByID[*ui.Button]("btnSystem"); err == nil {
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
