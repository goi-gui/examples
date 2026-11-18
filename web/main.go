//go:build js && wasm

// Command web runs the counter design in a browser: the same .goi file, tree,
// layout and painter as the desktop build, drawn with WebGPU into a canvas.
//
//	make -C ../examples web && python3 -m http.server -d ../examples/web/dist
package main

import (
	"embed"
	"fmt"
	"os"
	"strconv"

	"github.com/goi-gui/goi/ui"
)

//go:embed web.goi
var design embed.FS

// count lives outside Bind on purpose: Bind runs again on a hot reload, so
// state declared inside it would reset.
var count int

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fail(err)
	}
	// The page names the canvas the app draws into, so several apps can share
	// one page.
	win, err := app.LoadWindow("web.goi",
		ui.WithAssets(design), ui.WithSize(480, 400), ui.WithCanvas("goi-app"))
	if err != nil {
		fail(err)
	}
	err = win.Bind(func(w *ui.Window) error {
		label := w.MustByID[*ui.Text]("label")
		inc := w.MustByID[*ui.IconButton]("inc")
		dec := w.MustByID[*ui.IconButton]("dec")
		reset := w.MustByID[*ui.Button]("reset")

		render := func() {
			label.Text.Set(strconv.Itoa(count))
			dec.Disabled.Set(count <= 0)
			reset.Disabled.Set(count == 0)
		}
		render()

		inc.OnClick(func(ui.ClickEvent) {
			count++
			render()
		})
		dec.OnClick(func(ui.ClickEvent) {
			if count > 0 {
				count--
			}
			render()
		})
		reset.OnClick(func(ui.ClickEvent) {
			count = 0
			render()
		})
		return nil
	})
	if err != nil {
		fail(err)
	}
	if err := app.Run(); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	select {}
}
