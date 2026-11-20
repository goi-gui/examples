// Command shopping is a product page: a rail of categories, a picture column
// and the product's details beside it.
//
// Run with: make shopping (or cd shopping && goi run)
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

	err = app.LoadVariableFont("Nunito", "fonts/Nunito-VariableFont_wght.ttf", nil)
	if err != nil {
		panic(err)
	}

	err = app.LoadVariableFont("Nunito", "fonts/Nunito-Italic-VariableFont_wght.ttf", nil)
	if err != nil {
		panic(err)
	}

	win, err := app.LoadWindow("ui/Main.goi")
	if err != nil {
		panic(err)
	}

	images := [][]string{
		{
			"goi://images/model1pink.jpg",
			"goi://images/model2pink.png",
			"goi://images/model3pink.png",
		},
		{
			"goi://images/model1blue.jpg",
			"goi://images/model2blue.png",
			"goi://images/model3blue.png",
		},
		{
			"goi://images/model1purple.jpg",
			"goi://images/model2purple.png",
			"goi://images/model3purple.jpg",
		},
	}
	selectedImageIndex := 0

	isReviewsExpanded := true
	selectedSize := 6
	colors := []uint32{0xeac5c9ff, 0x7babe7ff, 0x432a80ff}
	selectedColorIndex := 0
	isFavourite := false

	setTheme := func(c uint32) {
		win.DefineTheme(ui.NewThemeFromSeed(c, false))
		win.DefineTheme(ui.NewThemeFromSeed(c, true))
	}

	setTheme(colors[selectedColorIndex])

	err = win.Bind(func(w *ui.Window) error {
		ui.SetData(w, "image1Path", images[selectedColorIndex][0])
		ui.SetData(w, "image2Path", images[selectedColorIndex][1])
		ui.SetData(w, "image3Path", images[selectedColorIndex][2])
		ui.SetData(w, "selectedImageIndex", selectedImageIndex)
		ui.SetData(w, "isReviewsExpanded", isReviewsExpanded)
		ui.SetData(w, "selectedSize", selectedSize)
		ui.SetData(w, "colors", colors)
		ui.SetData(w, "selectedColorIndex", selectedColorIndex)
		ui.SetData(w, "isFavourite", isFavourite)

		ui.SetHandler(w, "onToggleReviews", func() {
			isReviewsExpanded = !isReviewsExpanded
			ui.SetData(w, "isReviewsExpanded", isReviewsExpanded)
		})

		ui.SetHandler(w, "onToggleFavourite", func() {
			isFavourite = !isFavourite
			ui.SetData(w, "isFavourite", isFavourite)
		})

		ui.SetHandler(w, "onSelectColor", func(index int) {
			selectedColorIndex = index
			setTheme(colors[selectedColorIndex])
			ui.SetData(w, "selectedColorIndex", selectedColorIndex)
			ui.SetData(w, "image1Path", images[selectedColorIndex][0])
			ui.SetData(w, "image2Path", images[selectedColorIndex][1])
			ui.SetData(w, "image3Path", images[selectedColorIndex][2])
		})

		ui.SetHandler(w, "onSelectSize", func(size int) {
			selectedSize = size
			ui.SetData(w, "selectedSize", selectedSize)
		})

		ui.SetHandler(w, "onSelectImage", func(index int) {
			selectedImageIndex = index
			ui.SetData(w, "selectedImageIndex", selectedImageIndex)
		})
		return nil
	})
	if err != nil {
		panic(err)
	}

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
