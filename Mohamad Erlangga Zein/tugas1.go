package main
import "fmt"

func main() {
	//membuat variabel untuk menentukan tipe data
	var (
		satu, dua, tiga string
		temp string
	)

	//deklarasikan tipe data dengan masukkan fmt.Println dan fmt.Scanln
	fmt.Println("Masukkan input string: ")
	fmt.Scanln(&satu)
	fmt.Println("Masukkan input string: ")
	fmt.Scanln(&dua)
	fmt.Println("Masukkan input string: ")
	fmt.Scanln(&tiga)

	//output awal berisi kata yang akan diisi dan hasil dari input an nya berasal dari scanln nya
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	//gunakan var temp string untuk deklarasikan dari variabel satu, dua, tiga
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	//hasil output akhir akan berbeda karena sudah di deklarasi oleh temp sehingga urutan per kata nya berubah
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}