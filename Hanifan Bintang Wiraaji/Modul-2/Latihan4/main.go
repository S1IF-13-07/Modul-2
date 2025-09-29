package main

import "fmt"

func main() {
	var (
		c, f float64
	)

	fmt.Println("Konversi suhu Fahrenheit ke Celcius")
	fmt.Print("Masukan suhu Fahrenheit: ")
	fmt.Scanln(&f)
	c = ( f - 32 ) * 5 / 9
	fmt.Printf("Hasil konversi suhu %v Fahrenheit ke Celcius adalah %v", f, c)
}
