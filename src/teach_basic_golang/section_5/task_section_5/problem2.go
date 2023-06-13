package tasklist_section_5

import (
	"fmt"
)

// Fungsi untuk mencari angka-angka unik pada string input
func findUniqueNumbers(input string) []int {
	uniqueNumbers := []int{}
	countMap := make(map[int]int)

	// Menghitung jumlah kemunculan setiap angka
	for _, char := range input {
		number := int(char - '0')
		countMap[number]++
	}

	// Mencari angka-angka unik yang hanya muncul sekali
	for number, count := range countMap {
		if count == 1 {
			uniqueNumbers = append(uniqueNumbers, number)
		}
	}

	return uniqueNumbers
}

func AngkaUnique() {
	var input string

	fmt.Print("Masukkan Baris Pertama : ")
	fmt.Scanln(&input)
	// input := "76523752"
	uniqueNumbers := findUniqueNumbers(input)
	fmt.Println("Output:", uniqueNumbers)

	fmt.Print("Masukkan Baris Kedua : ")
	fmt.Scanln(&input)
	// input = "1122"

	uniqueNumbers = findUniqueNumbers(input)
	fmt.Println("Output:", uniqueNumbers)
}
