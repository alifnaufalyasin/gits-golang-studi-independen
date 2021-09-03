package main

import (
	"fmt"
)

func main() {
	// M Alif Naufal Yasin - Universitas Telkom
	var barang string
	var harga, quantity int
	fmt.Print("Nama Barang : ")
	fmt.Scanln(&barang)
	fmt.Print("Harga Barang : ")
	fmt.Scanln(&harga)
	fmt.Print("Quantity Barang : ")
	fmt.Scanln(&quantity)

	total := harga * quantity

	fmt.Printf("Barang %s berjumlah %d memiliki harga satuan %d dengan harga total %d", barang, quantity, harga, total)

}
