package main
import "fmt"

func main() {
	var f float64
	fmt.Print("Masukkan suhu Fahrenheit: ")
	fmt.Scanln(&f)
	c := (f - 32) * 5 / 9
	fmt.Printf("SUHU DALAM CELCIUS: %.1f\n", c)
}