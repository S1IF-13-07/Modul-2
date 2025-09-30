package main

import "fmt"

func main() {
	var (
		nama, nim, kelas string
	)

	fmt.Print("Masukkan Nama, NIM, dan Kelas: ")
	fmt.Scan(&nama, &nim, &kelas)
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}
