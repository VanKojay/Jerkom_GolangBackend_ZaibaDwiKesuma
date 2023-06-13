package main

import (
	"fmt"
	problem "section_5/task_section_5"
)

func main() {
	var taskList_section_5 float64

	fmt.Print("Pilih task yang ingin dicoba : ")
	fmt.Scanln(&taskList_section_5)

	switch taskList_section_5 {
	case 1:
		problem.ArrayGabungan()
	case 2:
		problem.ArrayGabungan_advc()
	case 3:
		problem.AngkaUnique()
	case 4:
		problem.PairWithTargetSum()
	default:
		problem.ArrayGabungan()
	}
}
