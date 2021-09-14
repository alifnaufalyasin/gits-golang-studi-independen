package main

import (
	"fmt"
	"math"
	"strconv"
)

// func deret(jumlah int, kali int, pangkat int, start int, max int, iter int) int {
// 	penjumlahan := start - 1 + jumlah
// 	perkalian := penjumlahan * kali
// 	pangkatan := int(math.Pow(float64(perkalian), float64(pangkat)))
// 	fmt.Print(pangkatan, " ")
// 	if iter < max {
// 		return start + deret(jumlah, kali, pangkat, start, max, iter+1)
// 	} else {

// 		return pangkatan
// 	}
// }

func deret(dikali bool, kali string, max int, iter int, pola int) int {
	pangkatan := int(math.Pow(float64(pola), float64(iter)))
	var perkalian int
	if dikali {
		if kali == "iter" {
			perkalian = pangkatan * (iter + 1)
		} else {
			angka, _ := strconv.Atoi(kali)
			perkalian = pangkatan * angka
		}
	} else {
		perkalian = pangkatan
	}
	if iter < max {
		fmt.Print(perkalian, " + ")

		return perkalian + deret(dikali, kali, max, iter+1, pola)
	} else {
		fmt.Print(perkalian)
		return perkalian
	}
}

func main() {
	var iterasi int
	fmt.Print("Input Max Iterasi : ")
	fmt.Scanln(&iterasi)
	fmt.Print("1. SUM = ")
	fmt.Println(" =", deret(false, "iter", iterasi, 0, 6))
	fmt.Print("2. SUM = ")
	fmt.Println(" =", deret(true, "5", iterasi, 0, 10))
	fmt.Print("3. SUM = ")
	fmt.Println(" =", deret(false, "iter", iterasi, 0, 7))
	fmt.Print("4. SUM = ")
	fmt.Println(" =", deret(true, "iter", iterasi, 0, 10))
	fmt.Print("5. SUM = ")
	fmt.Println(" =", deret(true, "5", iterasi, 0, 10))
}
