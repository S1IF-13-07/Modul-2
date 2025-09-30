package main
import "fmt"

func main() {
	var (
		nama, kelas, nim string
	)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Kelas: ")
	fmt.Scanln(&kelas)
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&nim)
	fmt.Println("Nama saya adalah", nama)
	fmt.Println("salah satu mahasiswa Prodi S1-IF dari kelas", kelas)
	fmt.Println("dengan NIM", nim)
}