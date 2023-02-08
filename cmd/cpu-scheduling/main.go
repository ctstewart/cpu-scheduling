package main

import (
	"time"

	"github.com/ctstewart/cpu-scheduling/internal/pkg/SJFAlgorithm"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// type process struct {
// 	id    int
// 	name  string
// 	burst int
// }

func main() {
	createUI()

	SJFAlgorithm.SJFAlgorithm()
	// var processes []process
	// for i := 0; i < 5; i++ {
	// 	p := process{i, fmt.Sprintf("P%d", i+1), rand.Intn(9) + 1}
	// 	processes = append(processes, p)
	// }
	// fmt.Println("Processes before SJF:")
	// printTableOfProcesses(processes)
	// sort.SliceStable(processes, func(i, j int) bool {
	// 	return processes[i].burst < processes[j].burst
	// })
	// fmt.Println("\nProcesses after SJF:")
	// printTableOfProcesses(processes)
}

// func printTableOfProcesses(processes []process) {
// 	fmt.Printf("%-15s %-15s\n", "Process", "Burst Time")
// 	for _, process := range processes {
// 		fmt.Printf("%-15s %-15d\n", process.name, process.burst)
// 	}
// }

func createUI() {
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
