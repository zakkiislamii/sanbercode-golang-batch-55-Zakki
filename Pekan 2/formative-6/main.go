package main

import (
	"fmt"
	"math"
)

// soal 1
func hitungLingkaran(luasLingkaran *float64, kelilingLingkaran *float64, jariJari float64) {
	*luasLingkaran = math.Pi * jariJari * jariJari
	*kelilingLingkaran = 2 * math.Pi * jariJari
}

// soal 2
func introduce(sentence *string, nama string, jenisKelamin string, pekerjaan string, umur string) {
	if jenisKelamin == "laki-laki" {
		panggilan := "Pak"
		*sentence = panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else {
		panggilan := "Bu"
		*sentence = panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
}

// soal 3
func tambahBuah(buahList *[]string, namaBuah string) {
	*buahList = append(*buahList, namaBuah)
}

// soal 4
func tambahDataFilm(title, duration, genre, year string, dataFilmList *[]map[string]string) {
	filmData := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilmList = append(*dataFilmList, filmData)
}

func main() {
	//soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64
	jariJari := 5.0
	hitungLingkaran(&luasLingkaran, &kelilingLingkaran, jariJari)
	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

	//soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println("\n" + sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{}
	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")
	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}

	//soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)
	// isi dengan jawaban anda untuk menampilkan data
	fmt.Println("\nData Film:")
	for i, film := range dataFilm {
		fmt.Printf("%d. ", i+1)
		for key, value := range film {
			fmt.Printf("%s: %s ", key, value)
		}
		fmt.Println()
	}
}
