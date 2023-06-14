package tasklist_section_6

import (
	"fmt"
)

// Mendefinisikan fungsi findCommonSubstring dengan dua parameter string, yaitu A dan B.
// Fungsi ini akan mencari dan mengembalikan substring yang sama antara A dan B.
func findCommonSubstring(A string, B string) string {

	// Mendeklarasikan variabel common untuk menyimpan substring yang sama antara A dan B.
	// Mendeklarasikan variabel minLength dengan panjang A sebagai nilai awalnya.
	// Jika panjang B lebih kecil dari minLength, maka minLength akan diubah menjadi panjang B.
	// Tujuan dari ini adalah agar kita hanya membandingkan karakter hingga mencapai akhir dari string terpendek.
	common := ""
	minLength := len(A)
	if len(B) < minLength {
		minLength = len(B)
	}
	// Melakukan perulangan untuk membandingkan karakter per karakter dari awal string hingga mencapai minLength.
	// Jika karakter pada indeks i dari A sama dengan karakter pada indeks i dari B, maka karakter tersebut akan ditambahkan ke common.
	// Jika terdapat perbedaan pada karakter, maka perulangan dihentikan dengan menggunakan break.
	for i := 0; i < minLength; i++ {
		if A[i] == B[i] {
			common += string(A[i])
		} else {
			break
		}
	}

	return common
}

func CompareString() {
	A := "AKA"
	B := "AKASHI"
	result := findCommonSubstring(A, B)
	fmt.Println("Output:", result)

	A = "KANGOORO"
	B = "KANG"
	result = findCommonSubstring(A, B)
	fmt.Println("Output:", result)
}
