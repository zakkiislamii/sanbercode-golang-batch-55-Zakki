package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"
	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int
	panjang_persegi_panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar_persegi_panjang, _ := strconv.Atoi(lebarPersegiPanjang)
	alas_segitiga, _ := strconv.Atoi(alasSegitiga)
	tinggi_segitiga, _ := strconv.Atoi(tinggiSegitiga)
	luasPersegiPanjang = panjang_persegi_panjang * lebar_persegi_panjang
	luasSegitiga = (alas_segitiga * tinggi_segitiga) / 2
	kelilingPersegiPanjang = (panjang_persegi_panjang * 2) + (2 * lebar_persegi_panjang)
	fmt.Println("\nJawaban soal 1")
	fmt.Println("Luas persegi panjang     :", luasPersegiPanjang)
	fmt.Println("Keliling persegi panjang :", kelilingPersegiPanjang)
	fmt.Println("Luas segitiga            :", luasSegitiga)

	//soal 2
	nilaiJohn := 80
	nilaiDoe := 50
	var indeksJohn, indeksDoe string
	if nilaiJohn >= 80 {
		indeksJohn = "A"
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		indeksJohn = "B"
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		indeksJohn = "C"
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		indeksJohn = "D"
	} else {
		indeksJohn = "E"
	}
	if nilaiDoe >= 80 {
		indeksDoe = "A"
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		indeksDoe = "B"
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		indeksDoe = "C"
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		indeksDoe = "D"
	} else {
		indeksDoe = "E"
	}
	fmt.Println("\nJawaban soal 2")
	fmt.Println("Indeks nilai John :", indeksJohn)
	fmt.Println("Indeks nilai Doe  :", indeksDoe)

	//soal 3
	tanggal := 4
	bulan := 3
	tahun := 2005
	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}
	fmt.Println("\nJawaban soal 3")
	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	//soal 4
	var generasi string
	if tahun >= 1944 && tahun <= 1964 {
		generasi = "Baby boomer"
	} else if tahun >= 1965 && tahun <= 1979 {
		generasi = "Generasi X"
	} else if tahun >= 1980 && tahun <= 1995 {
		generasi = "Generasi Y (Millenials)"
	} else if tahun >= 1995 && tahun <= 2015 {
		generasi = "Generasi Z"
	}
	fmt.Println("\nJawaban soal 4")
	fmt.Println(generasi)

}
