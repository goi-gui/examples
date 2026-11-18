// Command window shows the Window element's own properties: the size limits,
// the resize/minimize/maximize switches, fullscreen, and how a set from Go
// reaches the platform. The state props are reported back by the platform, so
// the status line follows the title bar buttons too.
//
// Run with: make window (or cd window && goi run)
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
		// The setters write the same cells the design's literals fill, so the
		// native window and the layout always agree. The tree keeps the values,
		// so a hot reload carries them over: no app state needed.
		setMin := func(a, b float64) {
			w.Root().SetMinWidth(a)
			w.Root().SetMaxWidth(b)
			render(w)
		}
		render(w)

		if sw, err := w.ByID[*ui.Switch]("lockResize"); err == nil {
			sw.OnToggle(func(on bool) {
				w.SetResizable(!on)
				render(w)
			})
		}
		if sw, err := w.ByID[*ui.Switch]("lockMin"); err == nil {
			sw.OnToggle(func(on bool) {
				w.SetMinimizable(!on)
				render(w)
			})
		}
		if sw, err := w.ByID[*ui.Switch]("lockMax"); err == nil {
			sw.OnToggle(func(on bool) {
				w.SetMaximizable(!on)
				render(w)
			})
		}
		if sw, err := w.ByID[*ui.Switch]("fs"); err == nil {
			sw.OnToggle(func(on bool) {
				w.SetFullscreen(on)
				render(w)
			})
		}
		if err := ui.SetHandler(w, "SetMin", func() { setMin(600, 900) }); err != nil {
			return err
		}
		if err := ui.SetHandler(w, "ClearLimits", func() { setMin(-1, -1) }); err != nil {
			return err
		}
		if err := ui.SetHandler(w, "Minimize", func() { w.Minimize() }); err != nil {
			return err
		}
		if err := ui.SetHandler(w, "Maximize", func() { w.SetMaximized(!w.Maximized()) }); err != nil {
			return err
		}
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

func render(w *ui.Window) {
	w.MustByID[*ui.Switch]("lockResize").Checked.Set(!w.Resizable())
	w.MustByID[*ui.Switch]("lockMin").Checked.Set(!w.Minimizable())
	w.MustByID[*ui.Switch]("lockMax").Checked.Set(!w.Maximizable())
	w.MustByID[*ui.Switch]("fs").Checked.Set(w.Fullscreen())
	w.MustByID[*ui.Text]("status").Text.Set(status(w))
	w.MustByID[*ui.Text]("state").Text.Set(state(w))
}

func status(w *ui.Window) string {
	resize := "resize on"
	if !w.Resizable() {
		resize = "resize off"
	}
	min := "minimize on"
	if !w.Minimizable() {
		min = "minimize off"
	}
	max := "maximize on"
	if !w.Maximizable() {
		max = "maximize off"
	}
	fs := "windowed"
	if w.Fullscreen() {
		fs = "fullscreen"
	}
	return resize + " · " + min + " · " + max + " · " + fs
}

// state reads the two props the platform reports back, plus the desktop's
// language, which the design also sees as `app.systemLanguage`.
func state(w *ui.Window) string {
	mode := "windowed"
	if w.Maximized() {
		mode = "maximized"
	}
	lang := w.SystemLanguage()
	if lang == "" {
		lang = "unknown"
	}
	return mode + " · language " + lang
}
