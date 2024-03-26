package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	kata1 := "Bootcamp"
	kata2 := " Digital"
	kata3 := " Skill"
	kata4 := " Sanbercode"
	kata5 := " Golang"
	fmt.Print("Jawaban soal 1: ")
	fmt.Println(kata1 + kata2 + kata3 + kata4 + kata5)

	//soal 2
	halo := "Halo Dunia"
	findDunia := "Dunia"
	replaceDunia := "Golang"
	fmt.Print("Jawaban soal 2: ")
	fmt.Println(strings.Replace(halo, findDunia, replaceDunia, 1))

	//soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	findkataKedua, findkataKetiga := "s", "r"
	replacekatakedua, replacekataketiga := "S", "R"
	perubahanKataKedua, perubahanKataKetiga := strings.Replace(kataKedua, findkataKedua, replacekatakedua, -1), strings.Replace(kataKetiga, findkataKetiga, replacekataketiga, -1)
	perubahanKataKeempat := strings.ToUpper(kataKeempat)
	fmt.Print("Jawaban soal 3: ")
	fmt.Println(kataPertama + " " + perubahanKataKedua + " " + perubahanKataKetiga + " " + perubahanKataKeempat)

	//soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"
	toString1, _ := strconv.Atoi(angkaPertama)
	toString2, _ := strconv.Atoi(angkaKedua)
	toString3, _ := strconv.Atoi(angkaKetiga)
	toString4, _ := strconv.Atoi(angkaKeempat)
	fmt.Print("Jawaban soal 4: ")
	fmt.Println(toString1 + toString2 + toString3 + toString4)

	//soal 5
	kalimat := "halo halo bandung"
	find := "halo"
	replaceWith := "Hi"
	kalimatReplace := strings.Replace(kalimat, find, replaceWith, -1)
	angka := 2021
	angkaString := strconv.Itoa(angka)
	fmt.Print("Jawaban soal 5: ")
	fmt.Println("\"" + kalimatReplace + "\" -" + angkaString)

}
