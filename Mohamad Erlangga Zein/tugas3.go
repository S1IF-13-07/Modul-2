package main
import "fmt"

func main(){
	const Phi = 3.14159
	var jariJari float64

	fmt.Print("Masukkan nilai jari jari: ")
	fmt.Scan(&jariJari)

	hasil := Phi * jariJari * jariJari
	fmt.Printf("Luas lingkaran: %.1f", hasil)
}