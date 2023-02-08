package UI

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func UI() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)

	// Make the clock update every tick
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	formattedTime := time.Now().Format("The time is 03:04:05")
	clock.SetText(formattedTime)
}
