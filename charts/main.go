// Charts: a line, an area, a bar, a pie and a scatter chart in one window.
//
// The design file styles the charts with properties (`grid`, `axis`, `legend`,
// `values`, `stacked`, `horizontal`, `smooth`, `points`, `innerRadius`,
// `startAngle`, `min`, `max`); the numbers come from Go through SetData.
//
// Hovering a point highlights it and shows a tooltip; clicking it fires onSelect
// with the point's index, which is what the status line under the charts shows.
//
// Run with: go run -ldflags=-linkmode=internal ./charts
package main

import (
	"fmt"
	"os"

	"github.com/goi-gui/goi/ui"
)

// chartData lives outside Bind: Bind runs again after a hot reload.
var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun"}

var revenue = ui.ChartData{
	Categories: months,
	Series: []ui.Series{
		{Label: "2025", Values: []float64{12, 18, 15, 24, 21, 29}},
		{Label: "2026", Values: []float64{16, 22, 19, 28, 26, 34}},
	},
}

var costs = ui.ChartData{
	Categories: months,
	Series: []ui.Series{
		{Label: "Hosting", Values: []float64{4, 5, 6, 6, 7, 8}},
		{Label: "People", Values: []float64{8, 9, 8, 10, 11, 11}},
	},
}

var channels = ui.ChartData{
	Categories: []string{"Direct", "Search", "Social", "Email"},
	Series: []ui.Series{
		{Values: []float64{42, 27, 18, 13}},
	},
}

var latency = ui.ChartData{
	Categories: months,
	Series: []ui.Series{
		{Label: "p50", Values: []float64{38, 35, 41, 32, 30, 28}},
	},
}

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	win, err := app.LoadWindow(ui.DirFS("charts"), "charts.goi",
		ui.WithHotReload(true), ui.WithSize(1100, 820))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := win.Bind(func(w *ui.Window) error {
		line, err := ui.ByID[*ui.LineChart](w, "revenue")
		if err != nil {
			return err
		}
		line.SetData(revenue)

		area, err := ui.ByID[*ui.AreaChart](w, "costs")
		if err != nil {
			return err
		}
		area.SetData(costs)

		bar, err := ui.ByID[*ui.BarChart](w, "bar")
		if err != nil {
			return err
		}
		bar.SetData(revenue)

		pie, err := ui.ByID[*ui.PieChart](w, "channels")
		if err != nil {
			return err
		}
		pie.SetData(channels)

		scatter, err := ui.ByID[*ui.ScatterChart](w, "latency")
		if err != nil {
			return err
		}
		scatter.SetData(latency)

		status, err := ui.ByID[*ui.Text](w, "status")
		if err != nil {
			return err
		}
		// One handler for every chart: the design passes the chart's name in and
		// the event carries the point index.
		pick := func(name string, data ui.ChartData) func(int) {
			return func(i int) {
				if i < 0 || i >= len(data.Categories) {
					status.Text.Set(name + ": nothing")
					return
				}
				value := 0.0
				if len(data.Series) > 0 && i < len(data.Series[0].Values) {
					value = data.Series[0].Values[i]
				}
				status.Text.Set(fmt.Sprintf("%s: %s = %g", name, data.Categories[i], value))
			}
		}
		line.OnSelect(pick("line", revenue))
		area.OnSelect(pick("area", costs))
		bar.OnSelect(pick("bar", revenue))
		pie.OnSelect(pick("pie", channels))
		scatter.OnSelect(pick("scatter", latency))
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
