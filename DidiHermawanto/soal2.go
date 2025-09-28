package main

import (
	"fmt"
)

func main() {
	var jumlah int
	var nama, nim, kelas string

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scanln(&jumlah)

	// Perulangan sesuai jumlah
	for i := 1; i <= jumlah; i++ {
		fmt.Println("\nData Mahasiswa ke-", i)

		// Input nama, nim, kelas
		fmt.Print("Masukkan Nama: ")
		fmt.Scanln(&nama)

		fmt.Print("Masukkan NIM: ")
		fmt.Scanln(&nim)

		fmt.Print("Masukkan Kelas: ")
		fmt.Scanln(&kelas)

		// Output resume
		fmt.Println("Perkenalkan saya adalah", nama,
			", salah satu mahasiswa Prodi S1-IF dari kelas", kelas,
			"dengan NIM", nim)
	}
}
