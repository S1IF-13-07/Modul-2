package main

import "fmt"

func main() {
	// Deklarasi variabel string
	// "satu", "dua", "tiga" untuk menampung input dari user
	// "temp" dipakai sebagai variabel sementara untuk menukar nilai
	var (
		satu, dua, tiga string
		temp            string
	)

	// Input string pertama dari user
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)

	// Input string kedua dari user
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	// Input string ketiga dari user
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	// Menampilkan hasil input awal sebelum ditukar
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	// Proses pertukaran nilai (shifting):
	// Nilai satu dipindahkan ke temp
	temp = satu
	// Nilai dua dipindahkan ke variabel satu
	satu = dua
	// Nilai tiga dipindahkan ke variabel dua
	dua = tiga
	// Nilai lama satu (yang disimpan di temp) dipindahkan ke variabel tiga
	tiga = temp

	// Menampilkan hasil akhir setelah pertukaran nilai
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
