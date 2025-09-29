package main

import "fmt"

func main() {
	var (
		nama, kelas string
		nim         int
	)

	fmt.Print("Masukan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan kelas: ")
	fmt.Scanln(&kelas)
	fmt.Print("Masukan NIM: ")
	fmt.Scanln(&nim)
	fmt.Printf("Halo, perkenalkan saya %s mahasiswa Telkom University kelas %s dengan NIM %d", nama, kelas, nim)
}
