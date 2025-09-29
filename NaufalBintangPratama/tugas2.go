package main

import "fmt"

func main() {
	var (
		nama string
		nim  string
		kelas string
	)

	// Input dari pengguna
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)

	fmt.Print("Masukkan kelas: ")
	fmt.Scanln(&kelas)

	// Output berupa resume singkat
	fmt.Println("Perkenalkan saya adalah", nama + ", salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM", nim + ".")
}