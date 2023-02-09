package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ctstewart/cpu-scheduling/internal/pkg/algorithms"
)

func Display() {
	a := app.New()
	w := a.NewWindow("CPU Scheduling Practice Quiz")
	w.Resize(fyne.NewSize(800, 400))

	// grid := container.New(layout.NewGridLayout(2), form(&qData, &aData), table(qData), table(aData))

	grid := container.New(layout.NewGridLayout(3))
	form := form(grid)
	grid.Add(form)
	grid.Add(table([][]string{{"Hello"}}))

	w.SetContent(grid)
	w.ShowAndRun()
}

func form(grid *fyne.Container) *widget.Form {
	algSelect := widget.NewSelect([]string{"SJF"}, func(s string) {})
	algSelect.SetSelected("SJF")

	difficultySelect := widget.NewSelect([]string{"Easy", "Medium", "Hard"}, func(s string) {})
	difficultySelect.SetSelected("Easy")

	var difficulty = map[string]int{
		"Easy":   0,
		"Medium": 1,
		"Hard":   2,
	}

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Algorithm", Widget: algSelect}, {Text: "Difficulty", Widget: difficultySelect}},
		OnSubmit: func() { // optional, handle form submission
			grid.Objects = grid.Objects[:1]
			qData, aData := algorithms.SJF(difficulty[difficultySelect.Selected])
			grid.Add(table(qData))
			grid.Add(widget.NewButton(
				"Show Answer",
				func() {
					grid.Objects = grid.Objects[:2]
					grid.Add(table(aData))
				},
			))
		},
	}

	return form
}

func table(data [][]string) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return table
}
