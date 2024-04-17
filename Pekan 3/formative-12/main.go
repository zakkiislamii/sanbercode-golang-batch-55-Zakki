package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiMahasiswa = []NilaiMahasiswa{}
var users = map[string]string{
	"zakki": "pw123",
}

func main() {
	http.HandleFunc("/nilai-mahasiswa", handleNilaiMahasiswa)
	http.ListenAndServe(":8080", nil)
}

func handleNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
		return
	}

	if pass, found := users[username]; !found || pass != password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid username or password")
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	var input NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding JSON: %v", err)
		return
	}

	defer r.Body.Close()

	if input.Nilai > 100 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Nilai maksimal adalah 100")
		return
	}

	var indeks string
	switch {
	case input.Nilai >= 80:
		indeks = "A"
	case input.Nilai >= 70:
		indeks = "B"
	case input.Nilai >= 60:
		indeks = "C"
	case input.Nilai >= 50:
		indeks = "D"
	default:
		indeks = "E"
	}

	input.ID = uint(len(nilaiMahasiswa) + 1)
	input.IndeksNilai = indeks
	nilaiMahasiswa = append(nilaiMahasiswa, input)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}
