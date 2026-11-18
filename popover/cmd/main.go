// Command popover shows the Popover component: three alignments, a panel that
// flips above the bottom edge instead of leaving the window, and a design that
// keeps one panel open through a bridge value.
//
// Run with: make popover (or cd popover && goi run)
package main

import (
	"fmt"
	"os"

	"github.com/goi-gui/goi/ui"
)

// open lives outside Bind on purpose: Bind runs again after every hot reload,
// so state declared inside it would reset on each save.
var open = "start"

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
		// A popover is always visible, so the design hides each panel with an
		// `if` on this value; clicking the anchor toggles it.
		if err := ui.SetHandler(w, "onOpen", func(which string) {
			if open == which {
				open = ""
			} else {
				open = which
			}
			ui.SetData(w, "open", open)
		}); err != nil {
			return err
		}
		ui.SetData(w, "open", open)
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
