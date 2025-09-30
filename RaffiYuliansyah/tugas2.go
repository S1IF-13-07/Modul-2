package main 

import "fmt"

func main(){
	var nama, kelas, nim string
	fmt.Print("Input Nama, NIM, Kelas : ")
	fmt.Scanf("%s %s %s",&nama, &nim, &kelas)
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s",nama , kelas, nim)
}