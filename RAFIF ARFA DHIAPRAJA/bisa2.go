package main

import "fmt"

func main() {
	var nama, kelas string
	var nim int

	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scan(&kelas)

	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s, dengan NIM %d", nama, kelas, nim)
}