package main

import (
	"fmt"

	"github.com/asvvvad/exchange"
)

func currency(angka int, mataUang string) string {
	ex := exchange.New("IDR")
	if mataUang == "Dolar" {
		usd, _ := ex.ConvertTo("USD", angka)
		hasil := fmt.Sprintf("$%.2f", usd)
		return hasil
	} else {
		eur, _ := ex.ConvertTo("EUR", angka)
		hasil := fmt.Sprintf("€%.2f", eur)
		return hasil
	}
}

func convertGBP(gbp int) (int, int, int, int) {
	var knuts int = gbp * 100
	var sickle int = knuts / 29
	var sisaKnuts int = knuts % sickle
	var galleon int = sickle / 17
	var sisaSickle int = sickle % galleon

	return knuts, galleon, sisaSickle, sisaKnuts
}

func main() {
	// M Alif Naufal Yasin - Universitas Telkom
	var pil int
	fmt.Println(`Silahkan pilih menu
1. Penukaran Rupiah ke Dollar
2. Penukaran Rupiah ke Euro
3. Penukaran GBP ke Knut, Sickle, Galleon ( Bonus )`)
	fmt.Scanln(&pil)
	switch pil {
	case 1:
		var rupiah int
		fmt.Print("Masukkan Rupiah : ")
		fmt.Scanln(&rupiah)
		hasil := currency(rupiah, "Dolar")
		fmt.Println("Hasil Konvert : " + hasil)

	case 2:
		var rupiah int
		fmt.Print("Masukkan Rupiah : ")
		fmt.Scanln(&rupiah)
		hasil := currency(rupiah, "Euro")
		fmt.Println("Hasil Konvert : " + hasil)

	case 3:
		var GBP int
		fmt.Print("Masukkan jumlah GB Pounds \t: ")
		fmt.Scanln(&GBP)
		knuts, galleon, sisaSickle, sisaKnuts := convertGBP(GBP)
		fmt.Println("Jumlah knut yang didapat \t=", knuts)
		fmt.Println("Hasil penukaran mendapatkan \t=", galleon, "Galleon(s)")
		fmt.Println("Sisa ditukar menjadi \t\t=", sisaSickle, "Sickle(s)")
		fmt.Println("Keping knut yang tersisa \t=", sisaKnuts, "Knut(s)")

	}

}
