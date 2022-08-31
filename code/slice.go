package main

import "fmt"

func main() {

	// tiga tipe data slice:
	// - pointer adalah penunjuk data pertama di array pada slice
	// - lenght adalah panjang dari slice
	// - capacity adalah kapasitas dari slice, dimana lenght tidak boleh lebih dari capacity

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)

	// function append
	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Joan")
	fmt.Println(slice3)

	// days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	// daySlice1 := days[5:]
	// daySlice1[0] = "Sabtu Baru"
	// daySlice1[1] = "Minggu Baru"
	// fmt.Println(days)

	// daySlice2 := append(daySlice1, "Libur Baru")
	// daySlice2[0] = "Ups"
	// fmt.Println(daySlice2)
	// fmt.Println(days)

}
