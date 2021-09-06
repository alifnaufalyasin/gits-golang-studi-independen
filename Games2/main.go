package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cariDeret(pilihan int, jumlah int) ([]string, int) {
	var bilDeret []string
	var total int = 0
	var i int = 1
	for len(bilDeret) < jumlah {
		// Ganjil
		if (pilihan == 1 || pilihan == 3) && i%2 == 1 {
			bilstr := strconv.Itoa(i)
			bilDeret = append(bilDeret, bilstr)
			total += i
		}

		// Genap
		if (pilihan == 2 || pilihan == 3) && i%2 == 0 {
			bilstr := strconv.Itoa(i)
			bilDeret = append(bilDeret, bilstr)
			total += i

		}
		i++
	}
	return bilDeret, total
}

func deretSpesial(jumlah int) []int {
	var bilDeret []int
	var i int = 1
	for len(bilDeret) < jumlah {
		// Ganjil
		if i%2 == 1 {
			bilDeret = append(bilDeret, i)
			// total += i
		}

		// Genap
		if i%2 == 0 {
			bilDeret = append(bilDeret, (i + 3))
			// total += (i + 3)

		}
		i++
	}
	return bilDeret
}

func main() {
	// M Alif Naufal Yasin - Universitas Telkom
	var pil int
	fmt.Print("Soal ke (1 atau 2) : ")
	fmt.Scanln(&pil)
	switch pil {
	case 1:
		var deret int
		fmt.Print("Total deret : ")
		fmt.Scanln(&deret)
		if deret > 0 {
			arrDeret, jumlah := cariDeret(2, deret)
			fmt.Println(arrDeret)
			fmt.Println(strings.Join(arrDeret[:], " + "), "=", jumlah)
		} else {
			fmt.Println("Tidak Boleh Negatif")
		}
	case 2:
		var deret int
		fmt.Print("Total deret : ")
		fmt.Scanln(&deret)
		if deret > 0 {
			arrDeret := deretSpesial(deret)
			fmt.Println(arrDeret)
		} else {
			fmt.Println("Tidak Boleh Negatif")
		}
	}

}
