package main

import "fmt"

func main() {
	//soal 1
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			if i%3 == 0 {
				fmt.Println(i, "- I Love Coding")
			} else {
				fmt.Println(i, "- Santai")
			}
		} else if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		}
	}
	fmt.Println()

	//soal 2
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
	fmt.Println()

	//soal 3
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var output []string
	for _, kata := range kalimat {
		if kata != "aku" && kata != "dan" {
			output = append(output, kata)
		}
	}
	fmt.Println(output)
	fmt.Println()

	//soal 4
	var sayuran = []string{}
	var tambahSayuran = append(sayuran,
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	)
	for t := 0; t < len(tambahSayuran); t++ {
		fmt.Printf("%d. %s\n", t+1, tambahSayuran[t])
	}
	fmt.Println()

	//soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}
	fmt.Println(satuan["panjang"] * satuan["lebar"] * satuan["tinggi"])
}
