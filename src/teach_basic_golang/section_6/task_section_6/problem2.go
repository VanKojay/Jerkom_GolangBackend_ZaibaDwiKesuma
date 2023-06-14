package tasklist_section_6

import (
	"fmt"
)

// Mendefinisikan fungsi shiftString yang menerima offset integer dan sebuah string sebagai parameter.
// Fungsi ini akan menghasilkan sebuah string baru dengan melakukan pergeseran huruf sesuai dengan offset.
func shiftString(offset int, str string) string {

	shifted := ""
	// Melakukan perulangan untuk setiap karakter dalam string str.
	// Variabel _ digunakan untuk mengabaikan indeks dalam perulangan, sedangkan char digunakan untuk menyimpan karakter saat ini.
	for _, char := range str {
		// Memeriksa apakah karakter saat ini adalah spasi.
		// Jika iya, maka karakter spasi ditambahkan ke shifted tanpa perubahan.
		// Jika tidak, maka menggunakan fungsi shiftLetter untuk menggeser karakter dan menambahkannya ke shifted.
		if char == ' ' {
			shifted += " "
		} else {
			shifted += string(shiftLetter(int(char), offset))
		}
	}
	// Mengembalikan nilai shifted sebagai string baru yang sudah mengalami pergeseran huruf.
	return shifted
}

// Mendefinisikan fungsi shiftLetter yang menerima dua parameter, yaitu letterCode dan offset.
// Fungsi ini akan menggeser huruf dengan kode yang diberikan sebanyak offset.
func shiftLetter(letterCode, offset int) rune {
	// Mendeklarasikan konstanta alphabetSize dengan nilai 26, yang merupakan jumlah huruf dalam alfabet.
	const alphabetSize = 26
	// Menghitung indeks baru untuk karakter yang digeser.
	// Mengurangi 97 dari letterCode untuk mengubah huruf ke kode ASCII (misalnya, 'a' menjadi 0, 'b' menjadi 1, dst.).
	// Menambahkan offset ke letterCode untuk melakukan pergeseran.
	// Menggunakan operasi modulo (%) dengan alphabetSize untuk menangani wrapping saat melebihi huruf 'z'.
	letterCode = (letterCode - 97 + offset) % alphabetSize
	// Memeriksa jika letterCode hasil pergeseran kurang dari 0, maka ditambahkan dengan alphabetSize untuk melakukan wrapping ke awal alfabet.
	if letterCode < 0 {
		letterCode += alphabetSize
	}
	return rune(letterCode + 97)
}

func CaesarChiper() {
	// Mendeklarasikan variabel offset dengan nilai 3 dan str dengan nilai "abc" sesuai dengan contoh test case pertama.
	// Memanggil fungsi shiftString dengan offset dan str sebagai argumen.
	// Menyimpan hasilnya dalam variabel result.
	// Mencetak hasil dengan menggunakan fmt.Println.
	offset := 3
	str := "abc"
	result := shiftString(offset, str)
	fmt.Println("Output:", result)

	// Mengubah nilai variabel offset menjadi 1 dan str menjadi "abcdefghijklmnopqrstuvwxyz" sesuai dengan contoh test case kedua.
	// Memanggil fungsi shiftString dengan offset dan str sebagai argumen.
	// Menyimpan hasilnya dalam variabel result.
	// Mencetak hasil dengan menggunakan fmt.Println.
	offset = 1
	str = "abcdefghijklmnopqrstuvwxyz"
	result = shiftString(offset, str)
	fmt.Println("Output:", result)

	// Mengubah nilai variabel offset menjadi 1000 dan str menjadi "abcdefghijklmnopqrstuvwxyz" sesuai dengan contoh test case ketiga.
	// Memanggil fungsi shiftString dengan offset dan str sebagai argumen.
	// Menyimpan hasilnya dalam variabel result.
	// Mencetak hasil dengan menggunakan fmt.Println.
	offset = 1000
	str = "abcdefghijklmnopqrstuvwxyz"
	result = shiftString(offset, str)
	fmt.Println("Output:", result)
}
