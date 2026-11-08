// Command gallery renders every Material 3 Expressive component in one window,
// one category page at a time.
// It is the visual reference for the component set and doubles as a manual
// test: click, type, drag and scroll.
//
// Run with: go run -ldflags=-linkmode=internal ./gallery
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goi-gui/goi/ui"
)

// page is one category: the panel it shows and the label used in the app bar and
// the navigation surfaces.
type page struct {
	id    string
	label string
}

var pages = []page{
	{"pageButtons", "Buttons"},
	{"pageInputs", "Inputs"},
	{"pageSurfaces", "Surfaces"},
	{"pageFeedback", "Feedback"},
	{"pageLayout", "Layout"},
	{"pageMotion", "Motion"},
	{"pageTable", "Table"},
	{"pageTypography", "Typography"},
	{"pageImage", "Image"},
}

// navIDs are the containers that own a destination selection: the rail and the
// bottom bar. The drawer items are wired one by one below, because they also
// close the drawer.
var navIDs = []string{"rail", "navbar"}

// galleryRow is one row of the gallery's for-loop demo.
type galleryRow struct {
	Title      string `goi:"title"`
	Supporting string `goi:"supporting"`
}

// repeatItem is one row of the gallery's Fill demo.
type repeatItem struct {
	label, icon string
}

// person is one row of the Table page, filled with ui.SetData.
type person struct {
	Name string `goi:"name"`
	Role string `goi:"role"`
	Team string `goi:"team"`
	Age  int    `goi:"age"`
}

// people lives outside Bind: Bind runs again after a hot reload, so the rows are
// app state and Bind only re-applies them.
var people = []person{
	{Name: "Kaan", Role: "Framework author", Team: "Core", Age: 31},
	{Name: "Ada", Role: "Mathematician", Team: "Research", Age: 36},
	{Name: "Linus", Role: "Kernel maintainer", Team: "Core", Age: 54},
	{Name: "Grace", Role: "Compiler engineer", Team: "Tooling", Age: 45},
	{Name: "Alan", Role: "Designer", Team: "Product", Age: 28},
	{Name: "Margaret", Role: "Flight software", Team: "Research", Age: 41},
	{Name: "Dennis", Role: "Systems programmer", Team: "Tooling", Age: 39},
	{Name: "Barbara", Role: "Language designer", Team: "Research", Age: 47},
	{Name: "Ken", Role: "Systems programmer", Team: "Core", Age: 52},
	{Name: "Radia", Role: "Network engineer", Team: "Product", Age: 33},
}

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Without Parse, flag.Arg never sees the path: the gallery would always open
	// its own design and ignore the argument.
	flag.Parse()
	path := "gallery/MainPage.goi"
	if arg := flag.Arg(0); arg != "" {
		path = arg
	}
	win, err := app.LoadWindow(ui.DirFS(filepath.Dir(path)), filepath.Base(path),
		ui.WithHotReload(true), ui.WithSize(1100, 900))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	app.ShowFPS(true)
	app.VSync(false)
	app.MaxFPS(120)

	// Which page is open is app state: a reload rebuilds the nodes, not the state.
	active := 0
	drawerOpen := false
	dark := false
	// Until someone uses a theme switch, `theme: app.systemMode` owns the
	// scheme. The first switch takes it over, and a reload keeps that choice.
	themeChosen := false

	if err := win.Bind(func(w *ui.Window) error {
		// An image that exists only in Go: SetImage takes any image.Image, so a
		// chart or a screenshot needs no file and no src.
		if gen, err := ui.ByID[*ui.Image](w, "generated"); err == nil {
			gen.SetImage(gradient(180, 120))
		}
		show := func(i int) {
			if i < 0 || i >= len(pages) {
				return
			}
			active = i
			for j, p := range pages {
				if n, err := ui.ByID[*ui.Column](w, p.id); err == nil {
					n.SetVisible(j == i)
				}
			}
			if bar, err := ui.ByID[*ui.TopAppBar](w, "appbar"); err == nil {
				bar.Label.Set(pages[i].label)
			}
			// Every nav surface shows the same selection, whichever one was used.
			for _, id := range navIDs {
				if h, _ := ui.ByID[*ui.NavigationRail](w, id); h != nil {
					selectChildren(h.Base(), i)
				}
				if h, _ := ui.ByID[*ui.NavigationBar](w, id); h != nil {
					selectChildren(h.Base(), i)
				}
			}
			for j, id := range []string{"drawerButtons", "drawerInputs", "drawerSurfaces", "drawerFeedback", "drawerLayout", "drawerMotion", "drawerTable", "drawerTypography", "drawerImage"} {
				if item, err := ui.ByID[*ui.NavigationItem](w, id); err == nil {
					item.M3().Selected.Set(j == i)
				}
			}
		}

		// The drawer slides in from the left edge and back out to it. translateX
		// is the animation (left reserves -1 for "auto") and visibility only
		// flips at the ends, so a half-open sheet keeps drawing; a second press
		// mid-slide continues from where it is instead of jumping.
		const drawerW = 320
		var drawerStop func()
		slideDrawer := func(frame *ui.Positioned, to float64, visibleAfter bool, done func()) {
			if drawerStop != nil {
				drawerStop()
				drawerStop = nil
			}
			from := frame.TranslateX.Get()
			if from == to {
				frame.SetVisible(visibleAfter)
				if done != nil {
					done()
				}
				return
			}
			ease := ui.Emphasized
			if to > from {
				ease = ui.Standard
			}
			drawerStop = app.Tween(250*time.Millisecond, ease, func(p float64) {
				frame.TranslateX.Set(from + (to-from)*p)
				if p >= 1 {
					frame.SetVisible(visibleAfter)
					if done != nil {
						done()
					}
				}
			})
		}
		setDrawer := func(open bool) {
			drawerOpen = open
			scrim, _ := ui.ByID[*ui.Rectangle](w, "scrim")
			frame, err := ui.ByID[*ui.Positioned](w, "drawerFrame")
			if err != nil {
				return
			}
			if open {
				if !frame.Base().Visible() {
					frame.TranslateX.Set(-drawerW)
					frame.SetVisible(true)
				}
				if scrim != nil {
					scrim.SetVisible(true)
				}
				slideDrawer(frame, 0, true, nil)
				return
			}
			if !frame.Base().Visible() {
				return
			}
			// The scrim stays for the slide out, then goes with the sheet.
			slideDrawer(frame, -drawerW, false, func() {
				if scrim != nil {
					scrim.SetVisible(false)
				}
			})
		}

		// Both theme controls drive this one function: the rail switch and the
		// one on the Layout page stay in step because neither owns the state.
		applyTheme := func(d bool) {
			themeChosen = true
			dark = d
			theme := "light"
			if d {
				theme = "dark"
			}
			w.Root().Theme.Set(theme)
			for _, id := range []string{"darkSwitch", "railTheme", "drawerTheme"} {
				if sw, err := ui.ByID[*ui.Switch](w, id); err == nil {
					sw.Checked.Set(d)
				}
			}
		}

		// Restore the state a reload must not lose. A still-unset theme is left
		// to `app.systemMode`, so the switch starts on whatever the desktop uses.
		show(active)
		setDrawer(drawerOpen)
		if themeChosen {
			applyTheme(dark)
		}

		// Navigation: a click on a destination reports its index.
		if rail, err := ui.ByID[*ui.NavigationRail](w, "rail"); err == nil {
			rail.OnSelect(show)
		}
		if bar, err := ui.ByID[*ui.NavigationBar](w, "navbar"); err == nil {
			bar.OnSelect(show)
		}
		for j, id := range []string{"drawerButtons", "drawerInputs", "drawerSurfaces", "drawerFeedback", "drawerLayout", "drawerMotion"} {
			idx := j
			if item, err := ui.ByID[*ui.NavigationItem](w, id); err == nil {
				item.OnClick(func(ui.ClickEvent) {
					show(idx)
					setDrawer(false)
				})
			}
		}

		// The app-bar leading icon opens the drawer. Its hit area is measured from
		// the bar's own edge, because a rail beside it moves the bar's left edge
		// off the window's.
		if bar, err := ui.ByID[*ui.TopAppBar](w, "appbar"); err == nil {
			bar.OnClick(func(e ui.ClickEvent) {
				local := e.X
				for n := ui.Node(bar); n != nil; n = n.Base().Parent() {
					local -= n.Base().Bounds().X
				}
				if local <= 56 && bar.Icon.Get() != "" {
					setDrawer(!drawerOpen)
				}
			})
		}
		if scrim, err := ui.ByID[*ui.Rectangle](w, "scrim"); err == nil {
			scrim.OnClick(func(ui.ClickEvent) { setDrawer(false) })
		}
		w.OnKey(func(e ui.KeyEvent) {
			if e.Key == ui.KeyEscape {
				setDrawer(false)
			}
		})

		// The data demo: the design's for loop builds one row per item, so
		// Bind only supplies the slice.
		ui.SetData(w, "rows", []galleryRow{
			{Title: "VPN will not connect", Supporting: "Ada"},
			{Title: "Printer queue is stuck", Supporting: "Grace"},
		})

		// The Table page: the same rows, and the Go side of sorting and
		// selection. SortBy takes the column key the design declared, so no
		// index is repeated in Go.
		ui.SetData(w, "people", people)
		if table, err := ui.ByID[*ui.Table](w, "tablePeople"); err == nil {
			showSorted := func(key string, ascending bool, cleared bool) {
				t, err := ui.ByID[*ui.Text](w, "tableSorted")
				if err != nil {
					return
				}
				if cleared {
					t.Text.Set("declaration order")
					return
				}
				dir := "ascending"
				if !ascending {
					dir = "descending"
				}
				t.Text.Set(fmt.Sprintf("sorted by %q, %s", key, dir))
			}
			showSelection := func() {
				t, err := ui.ByID[*ui.Text](w, "tableSelected")
				if err != nil {
					return
				}
				rows := table.SelectedRows()
				if len(rows) == 0 {
					t.Text.Set("no row selected")
					return
				}
				names := make([]string, 0, len(rows))
				for _, i := range rows {
					names = append(names, table.CellText(i, 0))
				}
				t.Text.Set(fmt.Sprintf("selected %d: %s", len(rows), strings.Join(names, ", ")))
			}
			table.OnSort(func(s ui.SortOrder) { showSorted(s.Key, s.Ascending, false) })
			table.OnSelect(func(int) { showSelection() })

			if b, err := ui.ByID[*ui.Button](w, "tableSortName"); err == nil {
				b.OnClick(func(ui.ClickEvent) {
					if table.SortBy("name", true) {
						showSorted("name", true, false)
					}
				})
			}
			if b, err := ui.ByID[*ui.Button](w, "tableSortAge"); err == nil {
				b.OnClick(func(ui.ClickEvent) {
					if table.SortBy("age", false) {
						showSorted("age", false, false)
					}
				})
			}
			if b, err := ui.ByID[*ui.Button](w, "tableClearSort"); err == nil {
				b.OnClick(func(ui.ClickEvent) {
					table.SortColumn.Set(-1)
					showSorted("", true, true)
				})
			}
			if b, err := ui.ByID[*ui.Button](w, "tableSelectFirstTwo"); err == nil {
				b.OnClick(func(ui.ClickEvent) {
					table.SetSelectedRows(0, 1)
					showSelection()
				})
			}
		}
		// The Layout page's repeat demo: one template row, three copies. A
		// reload rebuilds the tree, so this runs again from Bind and the rows
		// come back with the page; Fill replaces them instead of stacking.
		if list, err := ui.ByID[*ui.Column](w, "repeatList"); err == nil {
			items := []repeatItem{
				{label: "First copy", icon: "person"},
				{label: "Second copy", icon: "star"},
				{label: "Third copy", icon: "photo"},
			}
			if err := ui.Fill(w, list, items, func(row *ui.ListItem, it repeatItem) {
				row.Label.Set(it.label)
				row.Icon.Set(it.icon)
			}); err != nil {
				return err
			}
		}

		// The Motion page: the three dots travel the same distance and differ
		// only in easing. Stops are kept so a second Play cancels the first run
		// instead of leaving two tweens pulling the same Left property.
		var easingStops []func()
		playEasing := func() {
			for _, stop := range easingStops {
				stop()
			}
			easingStops = easingStops[:0]
			for _, d := range []struct {
				id   string
				ease ui.Easing
			}{
				{"dotLinear", ui.Linear},
				{"dotStandard", ui.Standard},
				{"dotEmphasized", ui.Emphasized},
			} {
				dot, err := ui.ByID[*ui.Positioned](w, d.id)
				if err != nil {
					continue
				}
				dot.Left.Set(0)
				easingStops = append(easingStops, app.Animate(&dot.Left, 264, 700*time.Millisecond, d.ease))
			}
		}
		if b, err := ui.ByID[*ui.Button](w, "playEasing"); err == nil {
			b.OnClick(func(ui.ClickEvent) { playEasing() })
		}

		// One value drives the slider and the progress ring, so the Motion page
		// shows that any numeric Prop can be animated.
		valueUp := false
		if b, err := ui.ByID[*ui.Button](w, "playValue"); err == nil {
			b.OnClick(func(ui.ClickEvent) {
				valueUp = !valueUp
				to := 0.9
				if !valueUp {
					to = 0.15
				}
				if sl, err := ui.ByID[*ui.Slider](w, "motionSlider"); err == nil {
					app.Animate(&sl.Value, to, 700*time.Millisecond, ui.Standard)
				}
				if pg, err := ui.ByID[*ui.CircularProgress](w, "motionProgress"); err == nil {
					app.Animate(&pg.Value, to, 900*time.Millisecond, ui.Emphasized)
				}
			})
		}

		// The rail switch and the one on the Layout page are both entry points to
		// the theme.
		for _, id := range []string{"railTheme", "darkSwitch", "drawerTheme"} {
			if sw, err := ui.ByID[*ui.Switch](w, id); err == nil {
				sw.OnToggle(applyTheme)
			}
		}

		// The animated Image and SVG demo: one switch drives all three clocks.
		if sw, err := ui.ByID[*ui.Switch](w, "animPause"); err == nil {
			sw.OnToggle(func(paused bool) {
				for _, id := range []string{"gifImage", "apngImage"} {
					if im, err := ui.ByID[*ui.Image](w, id); err == nil {
						im.Paused.Set(paused)
					}
				}
				if s, err := ui.ByID[*ui.SVG](w, "spinSVG"); err == nil {
					s.Paused.Set(paused)
				}
			})
		}

		// loop: false plays one pass and holds the last frame; the switch flips
		// the same property on every clip.
		if sw, err := ui.ByID[*ui.Switch](w, "animLoop"); err == nil {
			sw.OnToggle(func(loop bool) {
				for _, id := range []string{"gifImage", "apngImage"} {
					if im, err := ui.ByID[*ui.Image](w, id); err == nil {
						im.Loop.Set(loop)
					}
				}
				if s, err := ui.ByID[*ui.SVG](w, "spinSVG"); err == nil {
					s.Loop.Set(loop)
				}
			})
		}

		// Wire the stateful controls so the pages behave, not just draw.
		if sw, err := ui.ByID[*ui.Switch](w, "switchOn"); err == nil {
			sw.OnToggle(func(v bool) { fmt.Println("switch:", v) })
		}
		// The FAB menu opens and closes: a closed menu gives its actions no bounds,
		// so they neither draw nor take a click.
		if fab, err := ui.ByID[*ui.FABMenu](w, "fabMenu"); err == nil {
			fab.OnClick(func(ui.ClickEvent) {
				fab.Expanded.Set(!fab.Expanded.Get())
			})
		}
		if cb, err := ui.ByID[*ui.Checkbox](w, "check"); err == nil {
			cb.OnToggle(func(v bool) { fmt.Println("checkbox:", v) })
		}
		if sl, err := ui.ByID[*ui.Slider](w, "slider"); err == nil {
			sl.OnChangeValue(func(v float64) { fmt.Printf("slider: %.2f\n", v) })
		}
		if f, err := ui.ByID[*ui.TextField](w, "field"); err == nil {
			f.OnEdit(func(s string) { fmt.Println("field:", s) })
		}
		// The pickers own their value; a click on a cell or the dial sets it and
		// reports it, so the caption below shows what was picked.
		if dp, err := ui.ByID[*ui.DatePicker](w, "datePicker"); err == nil {
			dp.OnSubmit(func(v string) {
				if t, err := ui.ByID[*ui.Text](w, "dateValue"); err == nil {
					t.Text.Set("Picked " + v)
				}
			})
		}
		if tp, err := ui.ByID[*ui.TimePicker](w, "timePicker"); err == nil {
			tp.OnSubmit(func(v string) {
				if t, err := ui.ByID[*ui.Text](w, "timeValue"); err == nil {
					t.Text.Set("Picked " + v)
				}
			})
		}
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

// selectChildren marks destination i as selected. Only NavigationItems count as
// destinations, so the rail's theme switch below them does not shift the index.
func selectChildren(b *ui.NodeBase, i int) {
	n := 0
	for _, kid := range b.Children() {
		if kid.Kind() != "NavigationItem" {
			continue
		}
		c, ok := kid.(interface{ M3() *ui.Component })
		if !ok {
			continue
		}
		if m := c.M3(); m != nil {
			m.Selected.Set(n == i)
		}
		n++
	}
}

// gradient is the in-memory picture the gallery hands to Image.SetImage.
func gradient(w, h int) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			m.Set(x, y, color.RGBA{
				R: uint8(0x67 + x*80/w),
				G: uint8(0x50 + y*90/h),
				B: uint8(0xA4 - x*40/w),
				A: 0xFF,
			})
		}
	}
	return m
}
