package main

import "fmt"

func main() {
	nama := ""
	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&nama)
	fmt.Println("Selamat datang di Purwokerto City")
	fmt.Printf("Halo, %s", nama)
}
