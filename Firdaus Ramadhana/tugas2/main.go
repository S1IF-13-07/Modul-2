package main

import "fmt"

func main() {
	var nama, kelas string
	var nim int 
	fmt.Print("Masukkan Nama : ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Nim : ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan kelas : ")
	fmt.Scanln(&kelas)
	fmt.Println("=========================")
	fmt.Println("DATA DIRI MAHASIGMA TELYU")
	fmt.Println("=========================")
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi IF-07 dari kelas %s dengan NIM %d.\n",
		nama, kelas, nim)
}
