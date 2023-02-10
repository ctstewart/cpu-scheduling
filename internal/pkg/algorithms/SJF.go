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
	numProcesses := 5 + (difficulty * 2)
	calculationsLabels := []string{"TWT", "AWT", "TCT", "ACT"}
	calculations := []string{"", "", "", ""}

	for i := 0; i < numProcesses; i++ {
		p := process{i, fmt.Sprintf("P%d", i+1), rand.Intn(9) + 1 + (difficulty * 5)}
		processes = append(processes, p)
	}

	qData = [][]string{{"Process", "Burst Time"}}
	aData = make([][]string, len(qData))
	copy(aData, qData)

	for _, process := range processes {
		qData = append(qData, []string{process.name, strconv.Itoa(process.burst)})
	}

	for _, v := range calculationsLabels {
		qData = append(qData, []string{v, ""})
	}

	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].burst < processes[j].burst
	})

	for _, process := range processes {
		aData = append(aData, []string{process.name, strconv.Itoa(process.burst)})
	}

	var twt int
	for i := 0; i < len(processes)-1; i++ {
		processWaitTime := 0
		for j := 0; j <= i; j++ {
			processWaitTime += processes[j].burst
		}
		twt += processWaitTime
	}
	calculations[0] = strconv.Itoa(twt)

	awt := twt / len(processes)
	calculations[1] = strconv.Itoa(awt)

	var tct int
	for i := 0; i < len(processes); i++ {
		processWaitTime := 0
		for j := 0; j <= i; j++ {
			processWaitTime += processes[j].burst
		}
		tct += processWaitTime
	}
	calculations[2] = strconv.Itoa(tct)

	act := tct / len(processes)
	calculations[3] = strconv.Itoa(act)

	for i := 0; i < len(calculationsLabels); i++ {
		aData = append(aData, []string{calculationsLabels[i], calculations[i]})
	}

	return qData, aData
}
