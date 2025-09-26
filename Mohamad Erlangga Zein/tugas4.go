package main
import "fmt"

func main() {	
	var celcius, farenheit float64

	fmt.Print("Masukkan nilai farenheit: ")
	fmt.Scan(&farenheit)

	celcius = (farenheit-32) * 5/9
	fmt.Printf("Nilai konversi farenheit ke celciusnya adalah: %.1f", celcius)
}