package main
import "fmt"

func main() {
	var(
		nama string
		kelas string
		nim int
	)

	fmt.Print("Masukkan nama kamu: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan NIM kamu: ")
	fmt.Scan(&nim)

	fmt.Print("Dari kelas mana?: ")
	fmt.Scan(&kelas)

	fmt.Println("Perkenalkan nama saya", nama + ",", "salah satu mahasiswa prodi S1-IF dari kelas",
	kelas, "dengan NIM", nim)
}