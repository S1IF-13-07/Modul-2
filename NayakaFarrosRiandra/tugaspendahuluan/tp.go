package main

import "fmt"

func main() {
	fmt.Println("Selamat datang di telyu")
	nama := ""
	fmt.Print("Masukkan nama kalian :")
	fmt.Scan(&nama)
	fmt.Printf("Halo, %s", nama)
}
