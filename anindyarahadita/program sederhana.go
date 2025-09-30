package main
import "fmt"

func main() {
    var (
        nama string
    )
    fmt.Print("Masukkan Nama: ")
    fmt.Scanln(&nama)
    fmt.Println("Halo, nama saya adalah", nama)
}
