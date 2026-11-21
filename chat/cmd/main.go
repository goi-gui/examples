package main

import (
	"strings"
	"time"

	"github.com/goi-gui/goi/ui"
)

// Message is one row in the chat. The field tags are the names the design
// reads from the loop variable.
type Message struct {
	Author string `goi:"author"`
	Text   string `goi:"text"`
	Time   string `goi:"time"`
	Mine   bool   `goi:"mine"`
}

// incoming is what the timer appends. A chat gets its messages from somewhere;
// this keeps the sample self-contained.
var incoming = []Message{
	{Author: "Ada", Text: "The build is green again."},
	{Author: "Grace", Text: "Nice. I will take the review."},
	{Author: "Linus", Text: "Can someone look at the scroll edge case?"},
	{Author: "Ada", Text: "Already on it."},
}

func main() {
	app, err := ui.NewApp()
	if err != nil {
		panic(err)
	}
	defer app.Close()

	win, err := app.LoadWindow("ui/Main.goi")
	if err != nil {
		panic(err)
	}

	messages := []Message{
		{Author: "Ada", Text: "Hey! The chat example is ready.", Time: "09:41"},
		{Author: "You", Text: "Auto-scroll only when I am at the bottom, right?", Time: "09:42", Mine: true},
		{Author: "Ada", Text: "Exactly. Scroll up to read history and new messages wait.", Time: "09:43"},
		{Author: "Grace", Text: "The composer sends on Enter, and the send button reads the field.", Time: "09:44"},
		{Author: "You", Text: "And the jump button only appears when I leave the end?", Time: "09:45", Mine: true},
		{Author: "Linus", Text: "Yes. The list follows growth while progress is 1.", Time: "09:46"},
		{Author: "Ada", Text: "I left a longer note below to make the list scroll from the start.", Time: "09:47"},
		{Author: "Ada", Text: "Try scrolling up with the wheel, then wait for the next message.", Time: "09:48"},
	}
	atEnd := true

	err = win.Bind(func(w *ui.Window) error {
		chat := ui.MustByID[*ui.List](w, "chat")
		input := ui.MustByID[*ui.TextField](w, "input")
		chat.AutoScrollSmooth(true, 250*time.Millisecond)

		send := func(text string) {
			text = strings.TrimSpace(text)
			if text == "" {
				return
			}
			messages = append(messages, Message{
				Author: "You",
				Text:   text,
				Time:   time.Now().Format("15:04"),
				Mine:   true,
			})
			input.Text.Set("")
			ui.SetData(w, "messages", messages)
			// Sending is an explicit jump to the newest line, even when the
			// reader had scrolled away.
			chat.AutoScrollSmooth(true, 200*time.Millisecond)
			atEnd = true
			ui.SetData(w, "atEnd", true)
		}
		ui.SetHandler(w, "onSend", send)
		ui.SetHandler(w, "onSendCurrent", func() { send(input.Text.Get()) })
		ui.SetHandler(w, "onScroll", func(x, y, maxY, progress float64, byUser bool) {
			if !byUser {
				return
			}
			at := maxY <= 0 || progress >= 1
			if at != atEnd {
				atEnd = at
				ui.SetData(w, "atEnd", at)
			}
		})
		ui.SetHandler(w, "onAtEnd", func() {
			if !atEnd {
				atEnd = true
				ui.SetData(w, "atEnd", true)
			}
		})
		ui.SetHandler(w, "onJump", func() {
			chat.AutoScrollSmooth(true, 250*time.Millisecond)
			atEnd = true
			ui.SetData(w, "atEnd", true)
		})

		ui.SetData(w, "messages", messages)
		ui.SetData(w, "atEnd", atEnd)
		return nil
	})
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0; ; i++ {
			time.Sleep(4 * time.Second)
			app.Post(func() {
				m := incoming[i%len(incoming)]
				m.Time = time.Now().Format("15:04")
				messages = append(messages, m)
				ui.SetData(win, "messages", messages)
			})
		}
	}()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
