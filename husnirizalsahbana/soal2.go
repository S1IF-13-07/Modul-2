package main

import "fmt"

func main() {
	var nama, nim, kelas string

	// Input dari user
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas)

	// Output resume singkat
	fmt.Println("Perkenalkan saya adalah", nama,
		", salah satu mahasiswa Prodi S1-IF dari kelas", kelas,
		"dengan NIM", nim, ".")
}
