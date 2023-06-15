package tasklist_section_7

import (
	"fmt"
	"time"
)

func pow(x, n int) int {
	if n == 0 {
		return 1
	}

	result := 1
	for n > 0 {
		if n%2 != 0 {
			result *= x
		}
		x *= x
		n /= 2
	}

	return result
}

func FastExponation() {
	var x int
	var n int
	fmt.Print("Masukkan bilangan X : ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan bilangan X : ")
	fmt.Scanln(&n)
	start := time.Now()
	result := pow(x, n)

	fmt.Println(result)

	elapsed := time.Since(start)
	fmt.Println("Durasi waktu:", elapsed)
}
