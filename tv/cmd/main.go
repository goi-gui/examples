// Command tv shows a live television player: the channel list is bridge data
// filled from Go, and picking a row sets the URL the Video plays.
//
// Run with: make tv (or cd tv && goi run)
//
// The channels are public HLS streams and need network access.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/goi-gui/goi/ui"
)

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	win, err := app.LoadWindow("ui/Main.goi")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	isChannelListOpen := true

	err = win.Bind(func(w *ui.Window) error {
		ui.SetData(w, "channels", channels)
		ui.SetData(w, "selectedChannelIndex", -1)
		ui.SetData(w, "selectedChannelURL", "")
		ui.SetHandler(w, "onChannelPick", func(index int) {
			selectedChannel := channels[index]
			ui.SetData(w, "selectedChannelIndex", index)
			ui.SetData(w, "selectedChannelName", selectedChannel.Name)
			ui.SetData(w, "selectedChannelURL", selectedChannel.URL)
		})

		channelListBG := w.MustByID[*ui.Column]("channelListBG")

		ui.SetHandler(w, "toggleChannelList", func() {
			if isChannelListOpen {
				isChannelListOpen = !isChannelListOpen
				app.Tween(200*time.Millisecond, ui.Standard, func(f float64) {
					channelListBG.SetWidth(400 - (f * 400))
				})
				return
			}
			isChannelListOpen = !isChannelListOpen
			app.Tween(200*time.Millisecond, ui.Standard, func(f float64) {
				channelListBG.SetWidth(f * 400)
			})
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
