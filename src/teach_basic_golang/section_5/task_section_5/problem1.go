package tasklist_section_5

import (
	"fmt"
	"strings"
)

// Fungsi untuk memeriksa apakah sebuah elemen ada di dalam array
func contains(arr []string, element string) bool {
	for _, value := range arr {
		if value == element {
			return true
		}
	}

	return false
}

// Fungsi untuk menggabungkan dua array dan menghindari nama yang sama
func mergeArrays(arr1 []string, arr2 []string) []string {
	merged := []string{}

	// Menambahkan elemen dari arr1 ke dalam merged
	for _, value := range arr1 {
		lowercase := strings.ToLower(value)
		if !contains(merged, lowercase) {
			merged = append(merged, lowercase)
		}
	}

	// Menambahkan elemen dari arr2 ke dalam merged, jika belum ada di dalamnya
	for _, value := range arr2 {
		lowercase := strings.ToLower(value)
		if !contains(merged, lowercase) {
			merged = append(merged, lowercase)
		}
	}

	return merged
}

func ArrayGabungan() {
	// // Array pertama
	// arr1 := []string{"King", "Devil Jin", "Akuma"}

	// // Array kedua
	// arr2 := []string{"eddie", "steve", "geese"}

	// Array pertama
	arr1 := []string{"Sergei", "Jin"}

	// Array kedua
	arr2 := []string{"jin", "steve", "bryan"}

	// Menggabungkan dua array
	mergedArray := mergeArrays(arr1, arr2)

	// Menampilkan array yang telah digabungkan
	fmt.Println("Array yang telah digabungkan:", mergedArray)
}
