package main

import (
	"fmt"
	"math"
	"net/http"
)

func luasAlas(jarijari float64) float64 {
	return math.Pi * jarijari * jarijari
}

func kelilingAlas(jarijari float64) float64 {
	return 2 * math.Pi * jarijari
}

func volume(jarijari, tinggi float64) float64 {
	return math.Pi * jarijari * jarijari * tinggi
}

func main() {
	jarijari := 7.0
	tinggi := 10.0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "jariJari: %.2f, tinggi: %.2f, volume: %.2f, luas alas: %.2f, keliling alas: %.2f", jarijari, tinggi, volume(jarijari, tinggi), luasAlas(jarijari), kelilingAlas(jarijari))
	})
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
