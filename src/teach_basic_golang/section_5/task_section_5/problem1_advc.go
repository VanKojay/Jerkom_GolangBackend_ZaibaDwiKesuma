package tasklist_section_5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi untuk memeriksa apakah sebuah elemen ada di dalam array
func contains_advc(arr []string, element string) bool {
	for _, value := range arr {
		if value == element {
			return true
		}
	}
	return false
}

// Fungsi untuk menggabungkan dua array dan menghindari nama yang sama
func mergeArrays_advc(arr1 []string, arr2 []string) []string {
	merged := []string{}

	// Menambahkan elemen dari arr1 ke dalam merged
	for _, value := range arr1 {
		merged = append(merged, value)
	}

	// Menambahkan elemen dari arr2 ke dalam merged, jika belum ada di dalamnya
	for _, value := range arr2 {
		if !contains_advc(merged, value) {
			merged = append(merged, value)
		}
	}

	return merged
}

func ArrayGabungan_advc() {
	reader := bufio.NewReader(os.Stdin)

	// Membaca input pengguna untuk array pertama
	fmt.Println("Masukkan elemen-elemen array pertama (pisahkan dengan koma):")
	input1, _ := reader.ReadString('\n')

	// Memisahkan input pengguna menjadi elemen-elemen array pertama
	arr1 := strings.Split(strings.TrimSuffix(input1, "\n"), ",")

	// Membaca input pengguna untuk array kedua
	fmt.Println("Masukkan elemen-elemen array kedua (pisahkan dengan koma):")
	input2, _ := reader.ReadString('\n')

	// Memisahkan input pengguna menjadi elemen-elemen array kedua
	arr2 := strings.Split(strings.TrimSuffix(input2, "\n"), ",")

	// Menggabungkan dua array
	mergedArray_advc := mergeArrays_advc(arr1, arr2)

	// Menampilkan array yang telah digabungkan
	fmt.Println("Array yang telah digabungkan:", mergedArray_advc)
}
