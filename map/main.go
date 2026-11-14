// Command map searches PiriAPI for a place and draws its boundary on an
// OpenStreetMap map. It is the manual check for the Map component and a worked
// example of bridge: the design owns the search field, the result list and the
// handlers, and Go owns the two HTTP calls.
//
// Run with: go run -ldflags=-linkmode=internal ./map
//
// The free key is rate limited: one search and one place lookup per selection.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/goi-gui/goi/ui"
)

// PiriAPI keys are public client keys; this one is the example's free key.
const (
	apiBase = "https://piriapi.kaanksc.com/api/v1"
	apiKey  = "pk_29976cfdbd7c353ceb2ee62fe5fe711b075a813e4596f354638ddac58e7d5a50"
)

// Result is one row of the result list. The .goi file declares the same fields.
type Result struct {
	ID     string `goi:"id"`
	Label  string `goi:"label"`
	Detail string `goi:"detail"`
}

// App state lives outside Bind: Bind runs again after every hot reload, so a
// value declared inside it would be reset by an edit to the design.
var (
	results  []Result
	status   = "Insert a location name and press Enter."
	selected string
	// shown is the place on the map, kept so a hot reload puts the boundary and
	// the camera back instead of falling to the world view.
	shown *place
	// searchSeq and pickSeq drop the answer of a request the user has already
	// replaced by typing again or picking another row.
	searchSeq, pickSeq int
)

var client = &http.Client{Timeout: 20 * time.Second}

// place is the part of PiriAPI's place shape this example reads. Coordinates
// follow GeoJSON: center and bbox carry longitude first, geom nests polygon,
// ring, [lon, lat].
type place struct {
	ID          string          `json:"id"`
	AdminLevel  int             `json:"admin_level"`
	LocalName   string          `json:"localname"`
	CountryCode string          `json:"country_code"`
	Center      []float64       `json:"center"`
	BBox        []float64       `json:"bbox"`
	Geom        [][][][]float64 `json:"geom"`
}

type searchResponse struct {
	List []place `json:"list"`
}

func main() {
	app, err := ui.NewApp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	win, err := app.LoadWindow(ui.DirFS("map"), "map.goi",
		ui.WithHotReload(true), ui.WithSize(1100, 740))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = win.Bind(func(w *ui.Window) error {
		ui.SetData(w, "Results", results)
		ui.SetData(w, "Status", status)
		ui.SetData(w, "Selected", selected)
		ui.SetData(w, "Loading", false)
		if shown != nil {
			showPlace(w, *shown)
		}

		if err := ui.SetHandler(w, "onSearch", func(query string) {
			search(w, app, strings.TrimSpace(query))
		}); err != nil {
			return err
		}
		return ui.SetHandler(w, "onPick", func(id string) {
			pick(w, app, id)
		})
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

// search runs one full-text search and fills the list. The request itself runs
// off the UI goroutine; everything that touches a node runs in app.Post.
func search(w *ui.Window, app *ui.App, query string) {
	searchSeq++
	seq := searchSeq
	if query == "" {
		results, selected = nil, ""
		ui.SetData(w, "Results", nil)
		ui.SetData(w, "Selected", "")
		setStatus(w, "Insert a location name and press Enter.")
		world(w)
		return
	}
	ui.SetData(w, "Loading", true)
	setStatus(w, "Searching…")
	go func() {
		var out searchResponse
		err := get(apiBase+"/fts/search?limit=10&show-names=false&query="+url.QueryEscape(query), &out)
		app.Post(func() {
			if seq != searchSeq {
				return // a newer search owns the screen
			}
			ui.SetData(w, "Loading", false)
			if err != nil {
				results = nil
				ui.SetData(w, "Results", nil)
				setStatus(w, "Searching failed: "+err.Error())
				return
			}
			results = rows(out.List)
			ui.SetData(w, "Results", results)
			switch len(results) {
			case 0:
				setStatus(w, "No result found.")
			case 1:
				setStatus(w, "1 result. Select to see on map.")
			default:
				setStatus(w, fmt.Sprintf("%d results. Select to see on map.", len(results)))
			}
		})
	}()
}

// pick loads one place with its boundary and frames it on the map.
func pick(w *ui.Window, app *ui.App, id string) {
	pickSeq++
	seq := pickSeq
	selected = id
	ui.SetData(w, "Selected", id)
	setStatus(w, "Fetching boundaries…")
	go func() {
		var p place
		err := get(apiBase+"/places/"+url.PathEscape(id)+"?geom-detail=moderate&show-names=false", &p)
		app.Post(func() {
			if seq != pickSeq {
				return
			}
			if err != nil {
				setStatus(w, "Location can not loaded: "+err.Error())
				return
			}
			showPlace(w, p)
			setStatus(w, p.LocalName)
		})
	}()
}

// showPlace points the map at a place: its bbox frames it and its boundary is
// drawn from coordinates, so no GeoJSON round trip is needed.
func showPlace(w *ui.Window, p place) {
	shown = &p
	m := ui.MustByID[*ui.Map](w, "map")
	m.SetGeometry(geometry(p.Geom))
	if len(p.BBox) == 4 {
		m.BBox.Set(p.BBox)
	} else if len(p.Center) == 2 {
		m.BBox.Set(nil)
		m.SetView(p.Center[1], p.Center[0], 12)
	}
}

// world returns the map to the whole world when nothing is selected.
func world(w *ui.Window) {
	shown = nil
	m := ui.MustByID[*ui.Map](w, "map")
	m.BBox.Set(nil)
	m.SetGeometry(nil)
	m.SetView(0, 0, 2)
}

func setStatus(w *ui.Window, s string) {
	status = s
	ui.SetData(w, "Status", s)
}

// rows turns the API hits into the rows the design walks.
func rows(ps []place) []Result {
	out := make([]Result, 0, len(ps))
	for _, p := range ps {
		detail := strings.ToUpper(p.CountryCode)
		if p.AdminLevel > 0 {
			detail = fmt.Sprintf("%s · level %d", detail, p.AdminLevel)
		}
		out = append(out, Result{ID: p.ID, Label: p.LocalName, Detail: detail})
	}
	return out
}

// geometry converts PiriAPI's MultiPolygon coordinates into the map's geometry.
// The API writes [lon, lat]; ui.Point keeps the two apart by name.
func geometry(g [][][][]float64) *ui.MapGeometry {
	if len(g) == 0 {
		return nil
	}
	out := &ui.MapGeometry{}
	for _, poly := range g {
		rings := make([][]ui.Point, 0, len(poly))
		for _, ring := range poly {
			pts := make([]ui.Point, 0, len(ring))
			for _, c := range ring {
				if len(c) < 2 {
					continue
				}
				pts = append(pts, ui.Point{Lat: c[1], Lng: c[0]})
			}
			rings = append(rings, pts)
		}
		out.Polys = append(out.Polys, rings)
	}
	return out
}

// get runs one authenticated GET and decodes its JSON. A non-200 response turns
// PiriAPI's {"error": "..."} into the error, so the status line can show it.
func get(endpoint string, out any) error {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		var e struct {
			Error string `json:"error"`
		}
		if json.Unmarshal(body, &e) == nil && e.Error != "" {
			return fmt.Errorf("%s", e.Error)
		}
		return fmt.Errorf("%s", resp.Status)
	}
	return json.Unmarshal(body, out)
}
