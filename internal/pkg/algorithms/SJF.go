package algorithms

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type process struct {
	id    int
	name  string
	burst int
}

func SJF(difficulty int) (qData [][]string, aData [][]string) {
	var processes []process
	for i := 0; i < 5+(difficulty*2); i++ {
		p := process{i, fmt.Sprintf("P%d", i+1), rand.Intn(9) + 1 + (difficulty * 5)}
		processes = append(processes, p)
	}

	qData = [][]string{{"Process", "Burst Time"}}
	aData = make([][]string, len(qData))
	copy(aData, qData)

	for _, process := range processes {
		qData = append(qData, []string{process.name, strconv.Itoa(process.burst)})
	}

	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].burst < processes[j].burst
	})

	for _, process := range processes {
		aData = append(aData, []string{process.name, strconv.Itoa(process.burst)})
	}

	return qData, aData
}

func printTableOfProcesses(processes []process) {
	fmt.Printf("%-15s %-15s\n", "Process", "Burst Time")
	for _, process := range processes {
		fmt.Printf("%-15s %-15d\n", process.name, process.burst)
	}
}
