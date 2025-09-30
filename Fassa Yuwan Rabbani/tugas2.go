package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Nim: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas)
	fmt.Println("\n--- Resume Mahasiswa ---")
	fmt.Printf("Perkenalkan nama saya adalah %s, salah satu mahasiswa Prodi %s, dengan NIM yaitu %s", nama, kelas, nim)
}
