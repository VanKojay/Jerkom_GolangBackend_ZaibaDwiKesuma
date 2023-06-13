package main

import (
	"fmt"
	problem "section_6/task_section_6"
)

func main() {
	var taskList_section_6 float64

	fmt.Print("Pilih task yang ingin dicoba : ")
	fmt.Scanln(&taskList_section_6)

	switch taskList_section_6 {
	case 1:
		problem.Hello()

	default:
		problem.Hello()
	}
}
