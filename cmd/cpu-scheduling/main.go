package main

import (
	"github.com/ctstewart/cpu-scheduling/internal/pkg/UI"
	"github.com/ctstewart/cpu-scheduling/internal/pkg/algorithms"
)

// type process struct {
// 	id    int
// 	name  string
// 	burst int
// }

func main() {
	UI.Display()

	algorithms.SJF()
}
