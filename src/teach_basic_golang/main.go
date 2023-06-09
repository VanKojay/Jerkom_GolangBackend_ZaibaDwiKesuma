package main

import (
	"fmt"
	problem "teach_basic_golang/section3"
)

func main() {
	var task float64

	fmt.Print("Pilih task yang ingin dicoba : ")
	fmt.Scanln(&task)

	switch task {
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
