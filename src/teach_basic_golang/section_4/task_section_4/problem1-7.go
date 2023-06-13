package task_section_4

import (
	"fmt"
	"math"
)

func LuasPermukaanTabung() {
	var radius, height float64

	fmt.Print("Masukkan jari-jari tabung: ")
	fmt.Scanln(&radius)

	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scanln(&height)

	// Menghitung luas permukaan tabung
	baseArea := math.Pi * math.Pow(radius, 2)
	sideArea := 2 * math.Pi * radius * height
	totalArea := 2*baseArea + sideArea

	fmt.Printf("Luas permukaan tabung adalah %.2f\n", totalArea)
}

func StudentScore() {
	var Score int
	var grade string

	fmt.Print("Masukkan Nilai Mahasiswa: ")
	fmt.Scanln(&Score)

	switch {
	case Score >= 80 && Score <= 100:
		grade = "A"
	case Score >= 65 && Score < 79:
		grade = "B"
	case Score >= 50 && Score < 64:
		grade = "C"
	case Score >= 35 && Score < 49:
		grade = "D"
	case Score >= 0 && Score < 34:
		grade = "E"
	default:
		grade = "Masukkan Nilai dengan benar"
	}

	fmt.Println("Nilai Grade:", grade)
}

func FaktorBilangan() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	fmt.Printf("Faktor-faktor dari %d adalah: ", bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}

// fungsi menghitung bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Mengecek apakah bilangan memiliki faktor selain 1 dan dirinya sendiri
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {

			return false
		}
	}
	return true
}

func BilanganPrima() {

	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	if isPrime(bilangan) {
		fmt.Printf("%d adalah bilangan prima\n", bilangan)
	} else {
		fmt.Printf("%d bukan bilangan prima\n", bilangan)
	}

}

func isPalindrome(str string) bool {
	// Mengonversi string menjadi slice byte untuk mempermudah pemrosesan
	strBytes := []byte(str)

	// Inisialisasi indeks awal dan akhir
	i := 0
	j := len(strBytes) - 1

	// Mengecek karakter di kedua ujung secara berulang sampai indeks awal dan akhir bertemu
	for i < j {
		// Jika ada karakter yang tidak sama, bukan palindrom
		if strBytes[i] != strBytes[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func Palindrome() {

	var input string

	fmt.Print("Masukkan sebuah string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Printf("%s adalah palindrom\n", input)
	} else {
		fmt.Printf("%s bukan palindrom\n", input)
	}
}

func Asterisk() {
	var tinggi int

	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)

	for i := 1; i <= tinggi; i++ {
		// Mencetak spasi sebelum asterisk pada setiap baris
		for j := 1; j <= tinggi-i; j++ {
			fmt.Print(" ")
		}

		// Mencetak asterisk pada setiap baris
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}

		fmt.Println() // Pindah ke baris baru setelah mencetak baris
	}
}

func CetakTabelPerkalian() {
	var n int

	fmt.Print("Masukkan n: ")
	fmt.Scanln(&n)

	// Mencetak header kolom
	fmt.Print("   ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	// Mencetak garis pembatas
	fmt.Print("   ")
	for i := 1; i <= n; i++ {
		fmt.Print("----")
	}
	fmt.Println()

	// Mencetak tabel perkalian
	for i := 1; i <= n; i++ {
		fmt.Printf("%2d|", i)
		for j := 1; j <= n; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
