package sjfalgorithm

import (
	"fmt"
	"math/rand"
	"sort"
)

type process struct {
	id    int
	name  string
	burst int
}

func Algorithm() {
	var processes []process
	for i := 0; i < 5; i++ {
		p := process{i, fmt.Sprintf("P%d", i+1), rand.Intn(9) + 1}
		processes = append(processes, p)
	}
	fmt.Println("Processes before SJF:")
	printTableOfProcesses(processes)
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].burst < processes[j].burst
	})
	fmt.Println("\nProcesses after SJF:")
	printTableOfProcesses(processes)
}

func printTableOfProcesses(processes []process) {
	fmt.Printf("%-15s %-15s\n", "Process", "Burst Time")
	for _, process := range processes {
		fmt.Printf("%-15s %-15d\n", process.name, process.burst)
	}
}
