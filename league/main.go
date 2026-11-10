package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"os"

	"github.com/goi-gui/goi/ui"
)

type TeamRaw struct {
	Name            string `json:"name" goi:"name"`
	Position        int    `json:"position" goi:"position"`
	Played          int    `json:"played" goi:"played"`
	Won             int    `json:"won" goi:"won"`
	Draw            int    `json:"draw" goi:"draw"`
	Lost            int    `json:"lost" goi:"lost"`
	GoalsFor        int    `json:"goals_for" goi:"goalsFor"`
	GoalsAgainst    int    `json:"goals_against" goi:"goalsAgainst"`
	Points          int    `json:"points" goi:"points"`
	LastFiveResults string `json:"last_five_results" goi:"lastFiveResults"`
	Badge           string `json:"badge" goi:"badge"`
}

type Team struct {
	Name            string   `json:"name" goi:"name"`
	Position        int      `json:"position" goi:"position"`
	Played          int      `json:"played" goi:"played"`
	Won             int      `json:"won" goi:"won"`
	Draw            int      `json:"draw" goi:"draw"`
	Lost            int      `json:"lost" goi:"lost"`
	GoalsFor        int      `json:"goals_for" goi:"goalsFor"`
	GoalsAgainst    int      `json:"goals_against" goi:"goalsAgainst"`
	Points          int      `json:"points" goi:"points"`
	LastFiveResults []string `json:"last_five_results" goi:"lastFiveResults"`
	Badge           string   `json:"badge" goi:"badge"`
}

func main() {
	file, err := os.Open("./league/data/super_lig.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var teamsRaw []TeamRaw

	err = json.UnmarshalDecode(jsontext.NewDecoder(file), &teamsRaw)
	if err != nil {
		panic(err)
	}

	teams := make([]Team, 0, len(teamsRaw))

	for _, t := range teamsRaw {
		l5 := []string{}

		for _, l := range t.LastFiveResults {
			l5 = append(l5, string(l))
		}

		teams = append(teams, Team{
			Name:            t.Name,
			Position:        t.Position,
			Played:          t.Played,
			Won:             t.Won,
			Draw:            t.Draw,
			Lost:            t.Lost,
			GoalsFor:        t.GoalsFor,
			GoalsAgainst:    t.GoalsAgainst,
			Points:          t.Points,
			LastFiveResults: l5,
			Badge:           t.Badge,
		})
	}

	app, err := ui.NewApp()
	if err != nil {
		panic(err)
	}

	files := ui.DirFS("./league/ui/")

	win, err := app.LoadWindow(files, "league.goi", ui.WithHotReload(true))
	if err != nil {
		panic(err)
	}

	win.Bind(func(w *ui.Window) error {
		ui.SetData(w, "teams", teams)
		ui.SetHandler(w, "onToggleScheme", func(isDark bool) {
			if isDark {
				w.Root().Theme.Set("dark")
				return
			}
			w.Root().Theme.Set("light")
		})
		return nil
	})

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
