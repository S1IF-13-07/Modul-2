package main
import "fmt"

func main() {
	var (nama, nim, kelas string
	)
	fmt.Print("Masukkan Nama  : ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan NIM   : ")
	fmt.Scanln(&nim)

	fmt.Print("Masukkan Kelas : ")
	fmt.Scanln(&kelas)

	fmt.Println("Perkenalkan saya adalah " + nama + ",", "salah satu mahasiswa Prodi S1-IF dari kelas " + kelas + ",",
	 "dengan NIM " + nim + "." )
	fmt.Println("Saya hobi bernyanyi sehingga saat mandi pun saya bernyanyi. Sekian terima kasih.")
}
