//go:build js && wasm

// Command web runs a design in the browser: the same .goi file, tree, layout
// and painter as the desktop build, drawn with WebGPU into a canvas.
//
//	make -C ../examples web && python3 -m http.server -d ../examples/web/dist
package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/goi-gui/goi/ui"
)

//go:embed web.goi
var design embed.FS

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fail(err)
	}
	win, err := app.LoadWindow(design, "web.goi", ui.WithSize(560, 620))
	if err != nil {
		fail(err)
	}
	if err := win.Bind(func(w *ui.Window) error {
		count, err := ui.ByID[*ui.Text](w, "count")
		if err != nil {
			return err
		}
		trend, err := ui.ByID[*ui.LineChart](w, "trend")
		if err != nil {
			return err
		}
		steps := []float64{}
		draw := func() {
			trend.SetData(ui.ChartData{
				Categories: []string{"start", "clicks"},
				Series:     []ui.Series{{Label: "clicks", Values: []float64{0, float64(len(steps))}}},
			})
		}
		draw()
		return ui.SetHandler(w, "onStep", func(delta int) {
			if delta < 0 && len(steps) == 0 {
				return
			}
			steps = append(steps, float64(delta))
			count.Text.Set(fmt.Sprint(len(steps)))
			draw()
		})
	}); err != nil {
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
