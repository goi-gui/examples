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

//go:embed web.goi assets
var design embed.FS

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fail(err)
	}
	// The page can host several apps, each in its own canvas; this one names
	// the canvas it draws into.
	win, err := app.LoadWindow(design, "web.goi", ui.WithSize(640, 760), ui.WithCanvas("goi-app"))
	if err != nil {
		fail(err)
	}
	if err := win.Bind(func(w *ui.Window) error {
		status, err := ui.ByID[*ui.Text](w, "status")
		if err != nil {
			return err
		}
		name, err := ui.ByID[*ui.TextField](w, "name")
		if err != nil {
			return err
		}
		// A page picks files through the browser's own picker; the srcs it
		// returns read like any asset.
		pick, err := ui.ByID[*ui.Button](w, "pick")
		if err != nil {
			return err
		}
		pick.OnClick(func(ui.ClickEvent) {
			app.PickFiles(func(files []string) {
				if len(files) == 0 {
					status.Text.Set("nothing picked")
					return
				}
				status.Text.Set(fmt.Sprintf("picked %s", files[0]))
			}, ui.DialogFilters(ui.FileFilter{Name: "Images and video", Patterns: []string{"*.png", "*.jpg", "*.mp4"}}))
		})

		// The text field proves the keyboard and the browser's input method
		// reach the app; the label under the video shows what was typed.
		name.OnEdit(func(text string) { status.Text.Set(text) })

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
			status.Text.Set(fmt.Sprint(len(steps), " clicks"))
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
