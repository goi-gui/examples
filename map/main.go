// Command map draws OpenStreetMap tiles with a GeoJSON boundary over them. It
// is the manual check for the Map component: drag to pan, wheel to zoom.
//
// Run with: go run -ldflags=-linkmode=internal ./map
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
	if _, err := app.LoadWindowFile("map/map.goi", ui.WithHotReload(true)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
