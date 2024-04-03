package main

import (
	"fmt"
	"math"
	"strings"
)

// soal 1
type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

type persegiPanjang struct {
	panjang, lebar int
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return (2 * math.Pi * t.jariJari * t.tinggi) + (2 * math.Pi * t.jariJari * t.jariJari)
}

type balok struct {
	panjang, lebar, tinggi int
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang*b.lebar) + float64(b.panjang*b.tinggi) + float64(b.lebar*b.tinggi))
}

// soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type data_phone interface {
	data() string
}

func (p phone) data() string {
	return fmt.Sprintf("\nname: %s\nbrand: %s\nyear: %d\ncolors: %s\n", p.name, p.brand, p.year, strings.Join(p.colors, ", "))
}

// soal 3
func luasPersegi(sisi int, sisii bool) interface{} {
	if sisi == 0 {
		if sisii {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	} else {
		hasil := sisi * sisi
		if sisii {
			return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, hasil)
		} else {
			return hasil
		}
	}
}

func main() {
	// soal 1
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang
	bangunDatar = segitigaSamaSisi{alas: 10, tinggi: 5}
	fmt.Println("Luas segitiga sama sisi:", bangunDatar.luas())
	fmt.Println("Keliling segitiga sama sisi:", bangunDatar.keliling())
	bangunDatar = persegiPanjang{panjang: 10, lebar: 5}
	fmt.Println("Luas persegi panjang:", bangunDatar.luas())
	fmt.Println("Keliling persegi panjang:", bangunDatar.keliling())
	bangunRuang = tabung{jariJari: 10, tinggi: 5}
	fmt.Println("Volume tabung:", bangunRuang.volume())
	fmt.Println("Luas permukaan tabung:", bangunRuang.luasPermukaan())
	bangunRuang = balok{panjang: 10, lebar: 5, tinggi: 3}
	fmt.Println("Volume balok:", bangunRuang.volume())
	fmt.Println("Luas permukaan balok:", bangunRuang.luasPermukaan())

	// soal 2
	var myPhone data_phone = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	fmt.Println(myPhone.data())

	//soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//soal 4
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
