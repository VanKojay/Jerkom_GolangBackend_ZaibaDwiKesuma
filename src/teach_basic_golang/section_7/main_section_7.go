package main

import (
	"fmt"
	problem "section_7/task_section_7"
)

func main() {
	var taskList_section_7 float64

	fmt.Print("Pilih task yang ingin dicoba : ")
	fmt.Scanln(&taskList_section_7)

	switch taskList_section_7 {
	case 1:
		problem.Hello()

	default:
		problem.Hello()
	}
}
