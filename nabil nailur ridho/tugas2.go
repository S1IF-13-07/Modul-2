package main

import "fmt"

func main(){
	var nama, nim, kelas string
	
	fmt.Print("Masukan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan NIM: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukan Kelas: ")
	fmt.Scanln(&kelas)

	fmt.Print("Perkenalan saya adalah ", nama, ",salah satu mahasiswa prodi S1-IF dari kelas ", kelas, " dengan NIM ", nim)
}