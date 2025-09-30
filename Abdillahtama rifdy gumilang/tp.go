package main

import "fmt"

func main() {
	var nama string                                   //ini di gunakan untuk nyimpen data nama
	fmt.Print("masukan nama kalian: ")                //dan ini untuk inpt nama
	fmt.Scanln(&nama)                                 //ini buat ngambil variabel yang saya namain sebagai "nama" atau input itu sendiri
	fmt.Println("perkenalkan nama saya adalah", nama) //ini hasil akhir "nama"
}
