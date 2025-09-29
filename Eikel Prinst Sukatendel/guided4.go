package main

import "fmt"

func main() {
    var celcius float64
    var kelvin float64

    fmt.Print("Masukkan suhu dalam Celcius = ")
    fmt.Scan(&celcius)

    reamur := celcius * (4.0 / 5.0)
    fahrenheit := (celcius * (9.0 / 5.0)) + 32
    kelvin = celcius + 273.15 // Perbaikan rumus konversi Kelvin

    fmt.Printf("Temperatur Celcius = %.2f\n", celcius)
    fmt.Printf("Derajat Reamur = %.2f\n", reamur)
    fmt.Printf("Derajat Fahrenheit = %.2f\n", fahrenheit)
    fmt.Printf("Derajat Kelvin = %.2f\n", kelvin)
}
