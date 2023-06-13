package tasklist_section_5

import "fmt"

// Fungsi untuk mencari pasangan angka pada array yang jumlahnya sama dengan target
func findTwoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // Membuat map untuk menyimpan angka-angka yang telah dilihat sebelumnya

	for i, num := range nums {
		complement := target - num // Menghitung nilai komplemen dari angka saat ini
		if idx, ok := numMap[complement]; ok {
			return []int{idx, i} // Jika ditemukan pasangan angka yang jumlahnya sama dengan target, kembalikan indeks pasangan tersebut
		}
		numMap[num] = i // Tambahkan angka saat ini ke map dengan nilai indeks
	}

	return []int{} // Jika tidak ditemukan pasangan angka yang jumlahnya sama dengan target, kembalikan array kosong
}

func PairWithTargetSum() {
	nums := []int{1, 2, 3, 4, 6}
	target := 6
	result := findTwoSum(nums, target)
	fmt.Println("Output:", result)

	nums = []int{2, 5, 9, 11}
	target = 11
	result = findTwoSum(nums, target)
	fmt.Println("Output:", result)

	nums = []int{1, 3, 5, 7}
	target = 12
	result = findTwoSum(nums, target)
	fmt.Println("Output:", result)

	nums = []int{1, 4, 6, 8}
	target = 10
	result = findTwoSum(nums, target)
	fmt.Println("Output:", result)

	nums = []int{1, 5, 6, 7}
	target = 6
	result = findTwoSum(nums, target)
	fmt.Println("Output:", result)

}
