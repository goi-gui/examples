// Shapes: the primitive-quality test bed. Circles, rounded boxes with radii and
// borders, thick curved strokes and a stroked bezier path, all in one page.
//
// Run with: go run -ldflags=-linkmode=internal ./shapes
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/goi-gui/goi/ui"
)

func main() {
	msaa := flag.Int("msaa", 0, "multisample count: 1, 2, 4 or 8 (0 keeps the renderer's default)")
	flag.Parse()

	app, err := ui.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *msaa > 0 {
		app.MSAA(*msaa)
	}
	if _, err := app.LoadWindow(ui.DirFS("shapes"), "shapes.goi",
		ui.WithHotReload(true), ui.WithSize(760, 900),
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
