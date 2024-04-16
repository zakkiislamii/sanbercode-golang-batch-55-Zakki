package main

import (
	"fmt"
	"sort"
	"time"
)

// soal 1
func p(kalimat string, tahun int) {
	defer func() {
		fmt.Println("Ini defer")
	}()
	fmt.Println(kalimat)
	fmt.Println(tahun)
}

// soal 2
func kelilingSegitigaSamaSisi(sisi int, p bool) (string, error) {
	if sisi == 0 {
		if p {
			return "", fmt.Errorf("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			return "", fmt.Errorf("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}
	if p {
		keliling := 3 * sisi
		return fmt.Sprintf("Keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	} else {
		return fmt.Sprintf("%d", 3*sisi), nil
	}
}

// soal 3
func tambahAngka(nilai int, angka *int) {
	*angka += nilai
}
func cetakAngka(angka *int) {
	fmt.Println("Total angka:", *angka)
}

// soal 4

func tambahData(phones *[]string, data ...string) {
	*phones = append(*phones, data...)
}
func urutkanDanCetak(phones *[]string) {
	sort.Strings(*phones)
	for i, phone := range *phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//soal 1
	defer p("Golang Backend Development", 2021)
	fmt.Println()

	//soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
	fmt.Println()

	//soal 3
	angka := 1
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
	fmt.Println()

	//soal 4
	var phones = []string{}
	tambahData(&phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
	urutkanDanCetak(&phones)
}
