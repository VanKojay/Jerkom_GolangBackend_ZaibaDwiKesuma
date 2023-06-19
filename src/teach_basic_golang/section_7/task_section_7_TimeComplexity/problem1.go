package tasklist_section_7

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	// if n <= 1 {
	// 	return false
	// }

	// // Cek apakah bilangan dapat dibagi oleh angka 2 atau 3
	// if n == 2 || n == 3 {
	// 	return true
	// }
	// if n%2 == 0 || n%3 == 0 {
	// 	return false
	// }

	// Cek faktor-faktor bilangan hingga akar kuadrat dari n
	// karena faktor yang lebih besar dari akar kuadrat akan memiliki pasangan yang lebih kecil
	// sehingga tidak perlu diperiksa
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func BilanganPrima() {
	var num int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&num)

	start := time.Now()

	if isPrime(num) {
		fmt.Println("Bilangan prima")
	} else {
		fmt.Println("Bukan bilangan prima")
	}

	elapsed := time.Since(start)
	fmt.Println("Durasi waktu:", elapsed)
}
