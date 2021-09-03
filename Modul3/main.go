package main

import (
	"fmt"
	"strings"
)

func findDuplicate(word string) []string {
	var char []string
	var duplicate []string
	for i := 0; i < len(word); i++ {
		found := false
		for j := 0; j < len(char); j++ {
			if char[j] == string(word[i]) {
				found = true
			}
		}
		if found == true {
			dupFound := false
			for k := 0; k < len(duplicate); k++ {
				if duplicate[k] == string(word[i]) {
					dupFound = true
				}
			}
			if !dupFound {
				duplicate = append(duplicate, string(word[i]))
			}
		} else {
			char = append(char, string(word[i]))
		}
	}
	return duplicate
}

func hitungSuhu(input string, output string, suhu int) float32 {
	var res float32
	derajat := float32(suhu)
	switch input {
	case "c":
		if output == "f" {
			// Rumus Celscius ke Fahrenheit °F = °C × 1,8 + 32
			res = derajat*1.8 + 32
		} else if output == "k" {
			// Rumus Celscius ke Kelvin °K = °C + 273,15
			res = derajat + 273.15
		} else {
			res = derajat
		}

	case "f":
		if output == "c" {
			// Rumus Fahrenheit ke Celscius °C = (°F − 32) / 1,8
			res = (derajat - 32) / 1.8
		} else if output == "k" {
			// Rumus Fahrenheit ke Kelvin °K = (°F + 459,67) / 1,8
			res = (derajat + 459.67) / 1.8
		} else {
			res = derajat
		}

	case "k":
		if output == "c" {
			// Rumus Kelvin ke Celscius °C = °K − 273,15
			res = derajat - 273.15
		} else if output == "f" {
			// Rumus Kelvin ke Fahrenheit °F = °K × 1,8 − 459,67
			res = derajat*1.8 - 459.67
		} else {
			res = derajat
		}
	}
	return res
}

func deret(pilihan int, jumlah int) []int {
	var bilDeret []int
	var i int = 1
	for len(bilDeret) < jumlah {
		// Ganjil
		if (pilihan == 1 || pilihan == 3) && i%2 == 1 {
			bilDeret = append(bilDeret, i)
		}

		// Genap
		if (pilihan == 2 || pilihan == 3) && i%2 == 0 {
			bilDeret = append(bilDeret, i)
		}
		i++
	}
	return bilDeret
}

func currency(angka string) string {
	var kelipatan int = 0
	var arr []string
	for i := len(angka); i > 0; i-- {
		kelipatan++
		arr = append([]string{string(angka[i-1])}, arr...)
		if kelipatan%3 == 0 && i-1 != 0 {
			arr = append([]string{"."}, arr...)
		}
	}
	hasil := "Rp. " + strings.Join(arr[:], "")
	return hasil
}

func hitungLuasKeliling(p float32, l float32) string {
	luas := p * l
	keliling := 2*p + 2*l
	res := fmt.Sprintf("Luas : %.2f dan Keliling : %.2f", luas, keliling)

	return res
}

func main() {
	// M Alif Naufal Yasin - Universitas Telkom
	var pil int
	fmt.Println(`Silahkan pilih menu
1. program untuk mengetahui jumlah karakter yang diinput dan karakter yang sama. (A dan a dianggap beda)
2. program Konversi suhu ruangan antara celcius, fahrenheit dan kelvin.
3. program deret dengan jumlah deret sesuai dengan yang diinput
4. program input output berupa konversi dari integer(numerik) menjadi format currency (Rupiah).
5. program untuk menghitung Luas dan keliling dari persegi panjang.`)
	fmt.Scanln(&pil)
	switch pil {
	case 1:
		var kata string
		fmt.Print("Masukkan Kata : ")
		fmt.Scanln(&kata)
		duplicate := findDuplicate(kata)
		fmt.Printf("Jumlah karakter adalah %d \n", len(kata))
		fmt.Println("Karakter yang duplikat adalah", duplicate)
	case 2:
		var suhuInput, suhuOutput string
		var suhu int
		fmt.Print("Masukkan Suhu Input celcius (c), fahrenheit (f), kelvin (k) : ")
		fmt.Scanln(&suhuInput)
		fmt.Print("Masukkan Angka Suhu : ")
		fmt.Scanln(&suhu)
		fmt.Print("Masukkan Suhu Output celcius (c), fahrenheit (f), kelvin (k) : ")
		fmt.Scanln(&suhuOutput)
		hasil := hitungSuhu(suhuInput, suhuOutput, suhu)
		fmt.Printf("Hasil konversi = %d°%s sama dengan %.3f°%s", suhu, suhuInput, hasil, suhuOutput)
	case 3:
		var pilihan, jumlah int
		fmt.Print("Pilih deret bilangan  1.ganjil  2.genap  3.keduanya : ")
		fmt.Scanln(&pilihan)
		fmt.Print("Masukkan jumlah deret: ")
		fmt.Scanln(&jumlah)
		hasil := deret(pilihan, jumlah)
		fmt.Println("Hasil Bilangan Deret", hasil)
	case 4:
		var angka string
		fmt.Print("Masukkan Angka : ")
		fmt.Scanln(&angka)
		hasil := currency(angka)
		fmt.Println("Hasil Konversi Currency : ", hasil)
	case 5:
		var panjang, lebar float32
		fmt.Print("Masukkan Panjang : ")
		fmt.Scanln(&panjang)
		fmt.Print("Masukkan Lebar : ")
		fmt.Scanln(&lebar)
		hasil := hitungLuasKeliling(panjang, lebar)
		fmt.Println(hasil)
		// case 2:
		// case 2:

	}

}
