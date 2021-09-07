package main

import (
	"fmt"
)

func bikinSegitiga(terbalik bool, baris int) string {
	var hasil string = ""
	var kolom int = baris*2 - 1

	for i := 0; i < baris; i++ {
		for j := 1; j < kolom+1; j++ {
			if terbalik {
				if j <= i || j >= kolom+1-i {
					hasil += " "
				} else {
					hasil += "*"
				}
			} else {
				if j >= baris-i && j <= baris+i {
					hasil += "*"
				} else {
					hasil += " "
				}
			}

		}
		hasil += "\n"
	}
	return hasil
}

func main() {
	// M Alif Naufal Yasin - Universitas Telkom
	var jenis string
	var baris int
	fmt.Print("Masukkan Jenis Segitiga ? (Terbalik/Normal) ")
	fmt.Scanln(&jenis)
	fmt.Print("Masukkan Jumlah Baris ? ")
	fmt.Scanln(&baris)
	hasil := bikinSegitiga(jenis == "Terbalik", baris)
	fmt.Printf(hasil)
}
