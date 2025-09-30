package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Print("Masukkan Barisan Bilangan Penjumlahan: ")
	var hasil int
	fmt.Scan(&a, &b, &c, &d, &e)
	hasil = a + b + c + d + e
	fmt.Println("Hasil penjumlahan", a, b, c, d, e, "adalah", hasil)
}
