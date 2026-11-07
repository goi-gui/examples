// Table: a sortable data table the app fills from Go.
//
// The `bridge` block declares the rows and the handlers. Go sets the rows with
// ui.SetData; the `for` loop builds one TableRow per person. Clicking a sortable
// header sorts, clicking a row selects it, and Go can sort or select too.
//
// Run with: go run -ldflags=-linkmode=internal ./table
package main

import (
	"fmt"
	"os"

	"github.com/goi-gui/goi/ui"
)

type person struct {
	Name string `goi:"name"`
	Role string `goi:"role"`
	Team string `goi:"team"`
	Age  int    `goi:"age"`
}

// people lives outside Bind: Bind runs again after a hot reload, so the rows
// are app state and Bind only re-applies them.
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
	win, err := app.LoadWindow(ui.DirFS("table"), "table.goi",
		ui.WithHotReload(true), ui.WithSize(900, 760))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := win.Bind(func(w *ui.Window) error {
		ui.SetData(w, "people", people)
		ui.SetData(w, "count", len(people))

		table, err := ui.ByID[*ui.Table](w, "people")
		if err != nil {
			return err
		}
		// A header click hands Go the column key and the new direction.
		table.OnSort(func(s ui.SortOrder) {
			if t, err := ui.ByID[*ui.Text](w, "sorted"); err == nil {
				dir := "ascending"
				if !s.Ascending {
					dir = "descending"
				}
				t.Text.Set(fmt.Sprintf("sorted by %q, %s", s.Key, dir))
			}
		})
		// Row selection is a display index, so it always means the row the
		// reader sees, sorted or not.
		table.OnSelect(func(i int) {
			detail, err := ui.ByID[*ui.Text](w, "selected")
			if err != nil {
				return
			}
			if i < 0 || i >= len(table.TableRows()) {
				detail.Text.Set("no row selected")
				return
			}
			detail.Text.Set(fmt.Sprintf("selected: %s — %s (%s)",
				table.CellText(i, 0), table.CellText(i, 1), table.CellText(i, 2)))
		})

		// The buttons are the Go side of sorting: SortBy takes the column key
		// the design declared, so no index has to be repeated in Go.
		sortBy := func(key string, ascending bool) {
			if !table.SortBy(key, ascending) {
				return
			}
			if t, err := ui.ByID[*ui.Text](w, "sorted"); err == nil {
				dir := "ascending"
				if !ascending {
					dir = "descending"
				}
				t.Text.Set(fmt.Sprintf("sorted by %q, %s", key, dir))
			}
		}
		if b, err := ui.ByID[*ui.Button](w, "sortName"); err == nil {
			b.OnClick(func(ui.ClickEvent) { sortBy("name", true) })
		}
		if b, err := ui.ByID[*ui.Button](w, "sortAge"); err == nil {
			b.OnClick(func(ui.ClickEvent) { sortBy("age", false) })
		}
		if b, err := ui.ByID[*ui.Button](w, "clearSort"); err == nil {
			b.OnClick(func(ui.ClickEvent) {
				table.SortColumn.Set(-1)
				if t, err := ui.ByID[*ui.Text](w, "sorted"); err == nil {
					t.Text.Set("declaration order")
				}
			})
		}
		if b, err := ui.ByID[*ui.Button](w, "selectFirst"); err == nil {
			b.OnClick(func(ui.ClickEvent) {
				table.SetSelectedRow(0)
				table.Select.Emit(0)
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
