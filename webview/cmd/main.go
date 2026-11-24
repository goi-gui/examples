// Command webview embeds the system web view in a goi window: an embedded
// goi:// site with its own CSS and scripts, the page-to-Go message bridge, and
// Go calling back into the page.
//
// Run with: make webview (or cd webview && goi run). The webview build tag is
// in runFlags/buildFlags: the engine is included through auto when the design
// uses WebView, and always by all.
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
	defer app.Close()

	win, err := app.LoadWindow("ui/Main.goi")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = win.Bind(func(w *ui.Window) error {
		view := w.MustByID[*ui.WebView]("view")
		status := w.MustByID[*ui.Text]("status")

		// The page posts through the engine's channel; a message is a string
		// and carries no origin, so an app that shows third-party content must
		// treat it as untrusted input.
		view.OnMessage(func(msg string) {
			status.Text.Set("page -> Go: " + msg)
			fmt.Println("page:", msg)
		})

		w.MustByID[*ui.Button]("home").OnClick(func(ui.ClickEvent) {
			view.Load("goi://ui/site/index.html")
		})
		w.MustByID[*ui.Button]("site").OnClick(func(ui.ClickEvent) {
			view.Load("https://example.com")
		})
		w.MustByID[*ui.Button]("back").OnClick(func(ui.ClickEvent) {
			if !view.GoBack() {
				status.Text.Set("Go -> page: no history to go back to")
			}
		})
		w.MustByID[*ui.Button]("forward").OnClick(func(ui.ClickEvent) {
			if !view.GoForward() {
				status.Text.Set("Go -> page: no history to go forward to")
			}
		})
		w.MustByID[*ui.Button]("eval").OnClick(func(ui.ClickEvent) {
			// index.html defines demoBump; a plain site ignores the call.
			view.Eval("window.demoBump ? window.demoBump() : 0")
			status.Text.Set("Go -> page: demoBump()")
		})
		w.MustByID[*ui.Button]("reload").OnClick(func(ui.ClickEvent) {
			view.Reload()
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
