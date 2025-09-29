package main
import "fmt"
func main() {
	var r float64
	phi := 3.14
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)
	luas := phi * r * r
	fmt.Printf("Luas lingkaran adalah %.1f\n", luas)
}