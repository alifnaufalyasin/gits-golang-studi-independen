package main

import (
	"fmt"
	"strings"
)

func cekMenu(menu []string, pilihan string) (string, bool) {
	var ada bool = false
	for i := range menu {
		if strings.ToLower(menu[i]) == strings.ToLower(pilihan) {
			ada = true
		}
	}
	return pilihan, ada
}

func printMenu(menu []string) string {
	var hasil string = "Toko Makanan Indonesia\n===========================\n"
	for i := range menu {
		hasil += fmt.Sprintf("%s\n", menu[i])
	}
	hasil += "==========================="
	return hasil
}

func printPesanan(cart []string) string {
	var hasil string = ""
	for i := range cart {
		hasil += fmt.Sprintf("Pesanan anda : %s\n", cart[i])
	}
	hasil += "Terimakasih atas pesanannya"
	return hasil
}

func main() {
	var makanan []string = []string{"Tahu", "Tempe", "Sambal", "Nasi", "Lele", "Ayam"}
	var cart []string
	var pil string = "y"
	var menu string
	var ada bool
	fmt.Println(printMenu(makanan))
	for pil != "t" {
		fmt.Print("Masukkan menu pesanan anda dalam huruf ( eg: Tahu ) : ")
		fmt.Scanln(&menu)
		menu, ada = cekMenu(makanan, menu)
		if ada {
			cart = append(cart, menu)
			fmt.Print("Lanjutkan memesan (Y/T) ?  ")
			fmt.Scanln(&pil)
			pil = strings.ToLower(pil)
		} else {
			fmt.Println("Tidak ada makanan tersebut")
		}
	}
	fmt.Println(printPesanan(cart))
}
