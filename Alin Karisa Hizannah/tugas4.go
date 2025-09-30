package main

import "fmt"

func main() {
    var F, C int

    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scan(&F)

    // Konversi F -> C
    C = (F - 32) * 5 / 9

    fmt.Printf("Suhu dalam Celcius = %d\n", C)
}
