package main

import "fmt"

//soal 1
func luasPersegiPanjang(panjang int, lebar int) (luas int) {
	luas = panjang * lebar
	return
}

func kelilingPersegiPanjang(panjang int, lebar int) (keliling int) {
	keliling = 2 * (panjang + lebar)
	return
}

func volumeBalok(panjang int, lebar int, tinggi int) (volume int) {
	volume = panjang * lebar * tinggi
	return
}

//soal 2
func introduce(nama string, jenis_kelamin string, pekerjaan string, umur string) (kalimat string) {
	if jenis_kelamin == "laki-laki" {
		panggilan := "Pak"
		kalimat = panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else {
		panggilan := "Bu"
		kalimat = panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}
	return
}

//soal 3
func buahFavorit(nama string, buah ...string) string {
	var buahStr string
	for i, b := range buah {
		if i > 0 {
			buahStr += ", "
		}
		buahStr += "\"" + b + "\""
	}
	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", nama, buahStr)
}

//soal 4
func tambahDataFilm(dataFilm *[]map[string]string) func(string, string, string, string) {
	return func(title, jam, genre, tahun string) {
		*dataFilm = append(*dataFilm, map[string]string{
			"title": title,
			"jam":   jam,
			"genre": genre,
			"tahun": tahun,
		})
	}
}

func main() {
	//soal 1
	panjang := 12
	lebar := 4
	tinggi := 8
	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)
	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	//soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	//soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	//soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm := tambahDataFilm(&dataFilm)
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")
	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
