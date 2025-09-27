package main
import "fmt"

func main() {
	var nama string

	fmt.Print("Masukkan nama kamu: ")
	fmt.Scan(&nama)
	fmt.Printf("Hallo %s", nama)
}