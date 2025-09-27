package main
import "fmt"
func main() {
	var f float64
	fmt.Print("Masukan suhu Fahrenheit: ")
	fmt.Scan(&f)
	c := (f - 32) * 5 / 9
	fmt.Printf("SUHU DALAM CELCIUS: %.1f\n", c)
}