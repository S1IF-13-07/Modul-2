package main

import "fmt"

func main () {
	var nama, kelas, nim string
	fmt.Scanln(&nama, &kelas, &nim)
	fmt.Println("Halo perkenalkan saya", nama, "salah satu mahasiswa Prodi S1-IF dari kelas", kelas, "dengan NIM", nim)	 
}