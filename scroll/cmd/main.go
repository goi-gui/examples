// Command scroll shows the `scroll` property and the Go offset API.
//
// Run with: make scroll (or cd scroll && goi run)
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

	err = win.Bind(func(w *ui.Window) error {
		page := w.MustByID[*ui.Column]("page")

		w.MustByID[*ui.Button]("top").OnClick(func(ui.ClickEvent) {
			page.SetScrollOffset(0)
		})
		w.MustByID[*ui.Button]("end").OnClick(func(ui.ClickEvent) {
			// Any value past the end reads back as the clamped end, so the
			// handler needs no knowledge of the content height.
			page.SetScrollOffset(1e9)
			fmt.Println("offset:", page.ScrollOffset())
		})
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
