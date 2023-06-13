package tasklist_section_5

import "fmt"

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
		merged = append(merged, value)
	}

	// Menambahkan elemen dari arr2 ke dalam merged, jika belum ada di dalamnya
	for _, value := range arr2 {
		if !contains(merged, value) {
			merged = append(merged, value)
		}
	}

	return merged
}

func ArrayGabungan() {
	// Array pertama
	arr1 := []string{"King", "Devil Jin", "Akuma"}

	// Array kedua
	arr2 := []string{"eddie", "steve", "geese"}

	// Menggabungkan dua array
	mergedArray := mergeArrays(arr1, arr2)

	// Menampilkan array yang telah digabungkan
	fmt.Println("Array yang telah digabungkan:", mergedArray)
}
