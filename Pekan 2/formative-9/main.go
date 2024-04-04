package main

import (
	"fmt"
	"formative-9/soal8"
)

func main() {
	// Soal 1
	var bangunDatar soal8.HitungBangunDatar
	var bangunRuang soal8.HitungBangunRuang
	bangunDatar = soal8.SegitigaSamaSisi{Alas: 10, Tinggi: 5}
	fmt.Println("Luas segitiga sama sisi:", bangunDatar.Luas())
	fmt.Println("Keliling segitiga sama sisi:", bangunDatar.Keliling())
	bangunDatar = soal8.PersegiPanjang{Panjang: 10, Lebar: 5}
	fmt.Println("Luas persegi panjang:", bangunDatar.Luas())
	fmt.Println("Keliling persegi panjang:", bangunDatar.Keliling())
	bangunRuang = soal8.Tabung{JariJari: 10, Tinggi: 5}
	fmt.Println("Volume tabung:", bangunRuang.Volume())
	fmt.Println("Luas permukaan tabung:", bangunRuang.LuasPermukaan())
	bangunRuang = soal8.Balok{Panjang: 10, Lebar: 5, Tinggi: 3}
	fmt.Println("Volume balok:", bangunRuang.Volume())
	fmt.Println("Luas permukaan balok:", bangunRuang.LuasPermukaan())

	// Soal 2
	var myPhone soal8.DataPhone = soal8.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	fmt.Println(myPhone.Data())

	// Soal 3
	fmt.Println(soal8.LuasPersegi(4, true))
	fmt.Println(soal8.LuasPersegi(8, false))
	fmt.Println(soal8.LuasPersegi(0, true))
	fmt.Println(soal8.LuasPersegi(0, false))

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}
	numbers1, ok1 := kumpulanAngkaPertama.([]int)
	numbers2, ok2 := kumpulanAngkaKedua.([]int)
	if ok1 && ok2 {
		total := 0
		for _, num := range numbers1 {
			total += num
		}
		for _, num := range numbers2 {
			total += num
		}
		fmt.Printf("\n%s%d + %d + %d + %d = %d\n", prefix, numbers1[0], numbers1[1], numbers2[0], numbers2[1], total)
	} else {
		fmt.Println("Terjadi kesalahan dalam konversi tipe data.")
	}
}
