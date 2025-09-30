package main

import "fmt"

func main() {
    var nama, nim, kelas string

    fmt.Print("Masukkan Nama, NIM, dan Kelas: ")
    fmt.Scanln(&nama, &nim, &kelas)
    fmt.Println("Nama  :", nama)
    fmt.Println("NIM   :", nim)
    fmt.Println("Kelas :", kelas)

	fmt.Println("Perkenalkan saya adalah", nama, "salah satu mahasiswa prodi S1-IF dari kelas", kelas, "dengan NIM", nim,".")
}