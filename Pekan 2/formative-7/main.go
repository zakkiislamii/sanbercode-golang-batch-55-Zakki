package main

import "fmt"

//soal 1
type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

//soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

//soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(newColor string) {
	p.colors = append(p.colors, newColor)
}

//soal 4
type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	film := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilm = append(*dataFilm, film)
}

func main() {
	// soal 1
	buah1 := buah{"Nanas", "Kuning", false, 9000}
	buah2 := buah{"Jeruk", "Oranye", true, 8000}
	buah3 := buah{"Semangka", "Hijau & Merah", true, 10000}
	buah4 := buah{"Pisang", "Kuning", false, 5000}
	fmt.Println("Nama:", buah1.nama)
	fmt.Println("Warna:", buah1.warna)
	if buah1.adaBijinya {
		fmt.Println("Ada Bijinya: Ada")
	} else {
		fmt.Println("Ada Bijinya: Tidak")
	}
	fmt.Println("Harga:", buah1.harga)
	fmt.Println()
	fmt.Println("Nama:", buah2.nama)
	fmt.Println("Warna:", buah2.warna)
	if buah2.adaBijinya {
		fmt.Println("Ada Bijinya: Ada")
	} else {
		fmt.Println("Ada Bijinya: Tidak")
	}
	fmt.Println("Harga:", buah2.harga)
	fmt.Println()
	fmt.Println("Nama:", buah3.nama)
	fmt.Println("Warna:", buah3.warna)
	if buah3.adaBijinya {
		fmt.Println("Ada Bijinya: Ada")
	} else {
		fmt.Println("Ada Bijinya: Tidak")
	}
	fmt.Println("Harga:", buah3.harga)
	fmt.Println()
	fmt.Println("Nama:", buah4.nama)
	fmt.Println("Warna:", buah4.warna)
	if buah4.adaBijinya {
		fmt.Println("Ada Bijinya: Ada")
	} else {
		fmt.Println("Ada Bijinya: Tidak")
	}
	fmt.Println("Harga:", buah4.harga)

	//soal 2
	luasSegitiga := segitiga{4, 6}
	luasPersegi := persegi{4}
	luasPersegiPanjang := persegiPanjang{4, 3}
	fmt.Println()
	fmt.Println("Luas Segitiga:", ((luasSegitiga.alas * luasSegitiga.tinggi) / 2))
	fmt.Println("Luas Persegi:", (luasPersegi.sisi * luasPersegi.sisi))
	fmt.Println("Luas Persegi Panjang:", (luasPersegiPanjang.panjang * luasPersegiPanjang.lebar))

	//soal 3
	myPhone := phone{
		name:   "Galaxy S21",
		brand:  "Samsung",
		year:   2021,
		colors: []string{"Phantom Gray", "Phantom White"},
	}
	fmt.Println("\nNama HP:", myPhone.name)
	fmt.Println("Nama Brand:", myPhone.brand)
	fmt.Println("Tahun:", myPhone.year)
	fmt.Println("Warna sebelum penambahan:", myPhone.colors)
	myPhone.addColor("Phantom Violet")
	fmt.Println("Warna setelah penambahan:", myPhone.colors)

	//soal 4
	var dataFilm []movie
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)
	for i, film := range dataFilm {
		fmt.Printf("\n%d. title: %s\n", i+1, film.title)
		fmt.Printf("   duration: %d jam\n", film.duration/60)
		fmt.Printf("   genre: %s\n", film.genre)
		fmt.Printf("   year: %d\n", film.year)
	}
}
