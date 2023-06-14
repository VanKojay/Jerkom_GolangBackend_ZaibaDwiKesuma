package tasklist_section_6

import (
	"fmt"
)

func Hello() {
	// // 5 replace
	// s := "katak"

	// t := strings.Replace(s, "a", "o", -1)
	// // Baris ini menggantikan semua kemunculan substring "a" dengan substring "o" dalam string s.
	// // Parameter keempat dengan nilai -1 menandakan bahwa semua kemunculan harus digantikan.
	// fmt.Printf(t)

	// // 6 insert
	// p := "green"
	// index := 2
	// q := p[:index] + "HAY" + p[index:]
	// // Baris ini membuat variabel q dengan menggabungkan tiga bagian dari string p.
	// // Bagian pertama, p[:index], mengambil substring dari awal hingga sebelum indeks ke-2 (indeks 0 dan 1).
	// // Bagian kedua, "HAY", adalah string "HAY" yang ingin ditambahkan di tengah.
	// // Bagian ketiga, p[index:], mengambil substring dari indeks ke-2 hingga akhir string p.
	// fmt.Println(p, q)

	//Pointer
	// Jadi, dalam kode ini, kita mendeklarasikan variabel Name dengan nilai "John".
	// Kemudian, kita membuat variabel NameAddress yang merupakan pointer ke Name (menunjuk ke alamat memori Name).
	// Kemudian, kita mencetak nilai Name, alamat memori Name, nilai yang ditunjuk oleh pointer NameAddress (dengan dereference operator *),
	// dan nilai NameAddress.

	var Name string = "John"
	var NameAddress *string = &Name
	fmt.Println("Name Value : ", Name)
	fmt.Println("Name Addres Value : ", &Name)
	fmt.Println("Name Value : ", *NameAddress)
	fmt.Println("Name Addres Value : ", NameAddress)

	Name = "Doe"
	fmt.Println("Name Value : ", Name)
	fmt.Println("Name Addres Value : ", &Name)
	fmt.Println("Name Value : ", *NameAddress)
	fmt.Println("Name Addres Value : ", NameAddress)

}
