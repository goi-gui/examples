// Markdown: one document rendered by the Markdown component, from headings and
// emphasis through lists, quotes and code, in a scrolling column.
//
// Run with: go run -ldflags=-linkmode=internal ./markdown
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
