// SVG: one icon file recoloured by the design, by a theme role, and from Go.
//
// Run with: go run -ldflags=-linkmode=internal ./svg
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
	win, err := app.LoadWindowFile("svg/svg.goi", ui.WithHotReload(true))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// The roles the demo buttons pick, as literal colours so the choice is
	// visible without looking the scheme up.
	colors := []uint32{0x6750A4FF, 0x7D5260FF, 0xB3261EFF}
	names := []string{"primary", "tertiary", "error"}

	if err := win.Bind(func(w *ui.Window) error {
		icon := ui.MustByID[*ui.SVG](w, "live")
		label := ui.MustByID[*ui.Text](w, "liveColor")
		if err := ui.SetHandler(w, "Pick", func(index int) {
			icon.Color.Set(colors[index])
			label.Text.Set(names[index])
		}); err != nil {
			return err
		}

		// A document that exists only in Go: no file, no src, just the markup.
		// SetDataString and SetDataBytes differ only in the argument type.
		inline := ui.MustByID[*ui.SVG](w, "inline")
		docs := []string{
			`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"><path d="M12 2l2.9 6.3 6.9.8-5.1 4.7 1.4 6.8L12 17.3 5.9 20.6l1.4-6.8L2.2 9.1l6.9-.8z"/></svg>`,
			`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"><path d="M13 2L4 14h6l-1 8 9-12h-6z"/></svg>`,
		}
		draw := func(index int) {
			if index == 0 {
				inline.SetDataString(docs[0])
				return
			}
			inline.SetDataBytes([]byte(docs[1]))
		}
		draw(0)
		return ui.SetHandler(w, "Draw", draw)
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
