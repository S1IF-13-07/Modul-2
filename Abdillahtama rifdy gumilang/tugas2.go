package main

import "fmt"

func main() {
	var nama, nim, kelas string
	fmt.Printf("masukkan nama Tuan : ")
	fmt.Scanln(&nama)
	fmt.Printf("masukkan nimmu Tuan: ")
	fmt.Scanln(&nim)
	fmt.Printf("masukkan kelasmu tuan : ")
	fmt.Scanln(&kelas)
	fmt.Printf("Perkenalkan nama Baginda %s, salah satu mahasiswa prodi INFORMATIKA dari kelas %s dengan NIM %s\n", nama, kelas, nim)
}
