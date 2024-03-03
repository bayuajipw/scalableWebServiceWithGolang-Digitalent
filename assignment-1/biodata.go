package main

// // ============================ ASSIGNMENT 1 oleh Bayu Aji Putra Wibowo ============================

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama          string
	Alamat        string
	Pekerjaan     string
	AlasanMemilih string
}

var DaftarTeman = []Teman{
	{Nama: "Farrel", Alamat: "Jl. Sabar", Pekerjaan: "Developer", AlasanMemilih: "Ingin belajar bahasa pemrograman baru"},
	{Nama: "Zhofar", Alamat: "Jl. Cekatan", Pekerjaan: "Web", AlasanMemilih: "Meningkatkan keterampilan web service"},
	{Nama: "Fajar", Alamat: "Jl. Prihatin", Pekerjaan: "Data Scientist", AlasanMemilih: "Belajar selain menganalisis data"},
	{Nama: "Amry", Alamat: "Jl. Kasih", Pekerjaan: "Pertamina", AlasanMemilih: "Meningkatkan keterampilan teknologi dunia perminyakan"},
	{Nama: "Raihan", Alamat: "Jl. Sayang", Pekerjaan: "PLN", AlasanMemilih: "Meningkatkan keterampilan teknologi dunia kelistrikan"},
	{Nama: "Dean", Alamat: "Jl. Jalan", Pekerjaan: "Pegawai Swasta", AlasanMemilih: "Mengisi waktu senggang"},
	{Nama: "Dwiyan", Alamat: "Jl. Rindu", Pekerjaan: "Nelayan", AlasanMemilih: "Nelayan juga harus belajar teknologi"},
	{Nama: "Dede", Alamat: "Jl. Tabah", Pekerjaan: "Pegawai Bank", AlasanMemilih: "Meningkatkan keterampilan perbankan"},
	{Nama: "Hilal", Alamat: "Jl. Sesat", Pekerjaan: "Dokter", AlasanMemilih: "Gabut"},
	{Nama: "Saiful", Alamat: "Jl. Kafir", Pekerjaan: "Arsitek", AlasanMemilih: "Belajar saja"},
}

// Menampilkan informasi teman berdasarkan nomor absen
func InfoTeman(absen int) {
	if absen < 1 || absen > len(DaftarTeman) {
		fmt.Println("Nomor absen tidak valid. Masukkan nomor absen antara 1 dan", len(DaftarTeman))
		return
	}

	teman := DaftarTeman[absen-1]
	fmt.Printf("Informasi Teman Absen No.%v:\n", absen)
	fmt.Printf("Nama: %s\n", teman.Nama)
	fmt.Printf("Alamat: %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
	fmt.Printf("Alasan Memilih Kelas Golang: %s\n", teman.AlasanMemilih)
}

func main() {
	// Cek argumen nomor absen
	if len(os.Args) != 2 {
		fmt.Println("Gunakan: go run biodata.go <nomor_absen>")
		return
	}

	// Konversi argumen ke dalam bentuk integer
	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka.")
		return
	}

	// Menampilkan informasi teman berdasarkan nomor absen
	InfoTeman(absen)
}
