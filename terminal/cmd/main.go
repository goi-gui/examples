// Command terminal shows a real shell in a Terminal component. The app owns the
// pty and the process; the component only draws the grid and forwards input.
//
// Run with: make terminal (or cd terminal && goi run)
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/creack/pty"
	"github.com/goi-gui/goi/ui"
)

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer app.Close()

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}
	cmd := exec.Command(shell)
	cmd.Env = append(os.Environ(), "TERM=xterm-256color", "COLORTERM=truecolor")
	ptmx, err := pty.StartWithSize(cmd, &pty.Winsize{Rows: 24, Cols: 80})
	if err != nil {
		fmt.Fprintln(os.Stderr, "goi: start shell:", err)
		os.Exit(1)
	}
	defer func() { _ = ptmx.Close() }()

	win, err := app.LoadWindow("ui/Main.goi")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = win.Bind(func(w *ui.Window) error {
		term := ui.MustByID[*ui.Terminal](w, "term")
		term.SetWriter(ptmx)
		term.SetReader(ptmx)
		term.SetWaiter(func() (int, error) {
			err := cmd.Wait()
			if cmd.ProcessState == nil {
				return -1, err
			}
			return cmd.ProcessState.ExitCode(), err
		})
		term.OnResize(func(size ui.TerminalSize) {
			_ = pty.Setsize(ptmx, &pty.Winsize{
				Rows: uint16(size.Rows),
				Cols: uint16(size.Cols),
			})
		})
		term.OnExit(func(exit ui.TerminalExit) {
			if exit.Err != nil {
				fmt.Fprintln(os.Stderr, "goi: shell:", exit.Err)
			}
			app.Quit()
		})
		if label, err := w.ByID[*ui.Text]("shellLabel"); err == nil {
			label.Text.Set(filepath.Base(shell) + " in a pty")
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
