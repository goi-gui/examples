// Type: the M3 type scale, explicit sizes and weights, and text wrapping.
//
// Run with: go run -ldflags=-linkmode=internal ./typography
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
	if _, err := app.LoadWindowFile("typography/typography.goi", ui.WithHotReload(true), ui.WithSize(900, 780)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
