package main

import (
	"fmt"
	problem "section_4/task_section_4"
)

func main() {
	var taskList_section_4 float64

	fmt.Print("Pilih task yang ingin dicoba : ")
	fmt.Scanln(&taskList_section_4)

	switch taskList_section_4 {
	case 1:
		problem.LuasPermukaanTabung()
	case 2:
		problem.StudentScore()
	case 3:
		problem.FaktorBilangan()
	case 4:
		problem.BilanganPrima()
	case 5:
		problem.Palindrome()
	case 6:
		problem.Asterisk()
	case 7:
		problem.CetakTabelPerkalian()
	default:
		problem.LuasPermukaanTabung()
	}
}
