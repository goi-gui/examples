// Audio/Video: a player with its own controls, and a design that drives it.
//
// The media is decoded by ffmpeg behind the scenes, so the app needs ffmpeg on
// PATH (or GOI_FFMPEG pointing at it). The design says what to show: `controls`
// draws the play bar, `autoplay` starts on the first frame, and `loop`, `muted`,
// `volume` and `startAt` are the switches a player needs.
//
// The buttons below drive the same methods a design can call by id, so the
// same transport works from `.goi` and from Go.
//
// Run with: make media (or cd media && goi run)
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

	app.VSync(false)
	app.MaxFPS(120)
	app.ShowFPS(true)

	win, err := app.LoadWindow("ui/Main.goi")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := win.Bind(func(w *ui.Window) error {
		status, err := w.ByID[*ui.Text]("status")
		if err != nil {
			return err
		}
		video, err := w.ByID[*ui.Video]("video")
		if err != nil {
			return err
		}
		audio, err := w.ByID[*ui.Audio]("audio")
		if err != nil {
			return err
		}
		// Every button in the design calls the player by id, so there is no
		// transport to wire from Go; Go only hears the ends here.
		video.OnEnd(func(time.Duration) {
			status.Text.Set("the video finished")
		})
		audio.OnEnd(func(time.Duration) {
			status.Text.Set("the audio finished")
		})
		return nil
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
