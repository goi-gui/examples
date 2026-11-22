// Cursor: hover any tile to see the pointer shape its `cursor` names. The
// values are the CSS ones; a child inherits its parent's cursor until it
// declares its own, and `auto` returns to the framework's shape.
//
// Run with: make cursor (or cd cursor && goi run)
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

	// App state lives outside Bind, so a hot reload keeps the count.
	count := 0

	if err := win.Bind(func(w *ui.Window) error {
		// The button proves a cursor declaration does not swallow presses: the
		// tile still counts clicks like any other button.
		if b, err := w.ByID[*ui.Button]("clickMe"); err == nil {
			b.OnClick(func(ui.ClickEvent) {
				count++
				if n, err := w.ByID[*ui.Text]("clicks"); err == nil {
					n.Text.Set(fmt.Sprintf("%d clicks on the grab button", count))
				}
			})
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
