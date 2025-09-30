package main

import "fmt"

// Program ini berfungsi untuk memutar posisi tiga buah kata
// Kata pertama akan dipindah ke belakang, sementara kata kedua dan ketiga
// maju satu posisi ke depan. Mirip seperti antrian yang berputar.
func main() {
	// Kita butuh 4 variabel: 3 untuk menyimpan kata-kata yang diinput,
	// dan 1 variabel 'temp' sebagai tempat penyimpanan sementara
	var (
		satu, dua, tiga string
		temp            string
	)

	// Minta user untuk memasukkan kata pertama
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)

	// Minta user untuk memasukkan kata kedua
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)

	// Minta user untuk memasukkan kata ketiga
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	// Tampilkan dulu susunan kata sebelum diputar
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	// Sekarang kita mulai proses memutar posisi kata-kata:
	// Pertama, kita amankan dulu kata pertama ke temp (biar tidak hilang)
	temp = satu

	// Kedua, kita geser kata kedua ke posisi pertama
	satu = dua

	// Ketiga, kita geser kata ketiga ke posisi kedua
	dua = tiga

	// Terakhir, kita ambil kata pertama tadi dari temp dan taruh di posisi ketiga
	tiga = temp

	// Tampilkan hasil akhir setelah posisi kata-kata diputar
	// Contoh: "apel jeruk mangga" akan menjadi "jeruk mangga apel"
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
