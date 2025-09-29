package main

import "fmt"

func main() {
    var F float64 

    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scanln(&F)

    
    C := (F - 32) * 5 / 9

    fmt.Printf("Suhu %.1f Fahrenheit = %.1f Celcius\n", F, C)
}
