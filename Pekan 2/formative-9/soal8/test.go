package soal8

import (
	"fmt"
	"math"
	"strings"
)

// Soal 1
type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return (2 * math.Pi * t.JariJari * t.Tinggi) + (2 * math.Pi * t.JariJari * t.JariJari)
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * (float64(b.Panjang*b.Lebar) + float64(b.Panjang*b.Tinggi) + float64(b.Lebar*b.Tinggi))
}

// Soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type DataPhone interface {
	Data() string
}

func (p Phone) Data() string {
	return fmt.Sprintf("\nname: %s\nbrand: %s\nyear: %d\ncolors: %s\n", p.Name, p.Brand, p.Year, strings.Join(p.Colors, ", "))
}

// Soal 3
func LuasPersegi(sisi int, sisii bool) interface{} {
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
