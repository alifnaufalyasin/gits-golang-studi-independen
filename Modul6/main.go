package main

import (
	"fmt"
	"sort"
)

type game struct {
	judul  string
	rating float32
	rate   string
}

func addData(data []game, judul string, rate float32) []game {
	var temp game
	temp.judul = judul
	if rate > 5.0 {
		rate = 5.0
	}
	if rate < 0.0 {
		rate = 0.0
	}
	temp.rating = rate
	if rate > 4.0 {
		temp.rate = "good"
	} else if rate > 2.0 {
		temp.rate = "average"
	} else {
		temp.rate = "poor"
	}

	data = append(data, temp)

	sort.Slice(data, func(i, j int) bool {
		return data[i].rating > data[j].rating
	})

	return data
}

func deleteData(data []game, index int) []game {
	return append(data[:index-1], data[index:]...)

}

func viewAllData(data []game, title string) string {
	var hasil string = title
	for i := 0; i < len(data); i++ {
		hasil += fmt.Sprintf("%d. %s rating: %.1f (%s) \n", i+1, data[i].judul, data[i].rating, data[i].rate)
	}
	hasil += "\n"
	return hasil
}

func searchByJudul(data []game, judul string) string {
	var found bool = false
	var hasil string
	for i := range data {
		if data[i].judul == judul {
			found = true
			hasil = fmt.Sprintf("%d. %s rating: %.1f (%s)", i+1, data[i].judul, data[i].rating, data[i].rate)
			break
		}
	}
	if found {
		return hasil
	} else {
		return "Data Tidak Ditemukan"
	}
}

func viewDataByRate(data []game, rate float32) string {
	var hasil string = fmt.Sprintf("List Data dengan Rating diatas %.1f\n", rate)
	for i := 0; i < len(data); i++ {
		if data[i].rating >= rate {
			hasil += fmt.Sprintf("%d. %s rating: %.1f (%s) \n", i+1, data[i].judul, data[i].rating, data[i].rate)
		}
	}
	hasil += "\n"
	return hasil
}

func main() {
	var pilihan int = 1
	var data []game
	for pilihan != 0 {
		fmt.Println(`Menu
1. Input Data Game Baru
2. Hapus Data Game berdasarkan ID Game
3. View Data Game beserta jumlah data yang tersimpan
4. Cari Data Game berdasarkan Nama
5. Top 3 Game Terfavorit
6. View Data dengan rating diatas yang diinputkan
0. Exit`)
		fmt.Print("Silahkan Pilih Menu : ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			var judul string
			var rating float32
			fmt.Print("Masukkan Judul Game : ")
			fmt.Scanln(&judul)
			fmt.Print("Masukkan Rating Game (0.0 - 5.0) : ")
			fmt.Scanln(&rating)
			data = addData(data, judul, rating)
			fmt.Println(data)
		case 2:
			fmt.Println(viewAllData(data, "List Game \n"))
			var idx int
			fmt.Print("Nomor data yang ingin dihapus : ")
			fmt.Scanln(&idx)
			data = deleteData(data, idx)
			fmt.Println(viewAllData(data, "List Game \n"))
		case 3:
			fmt.Print(viewAllData(data, "List Game \n"))
			fmt.Println("Total Data :", len(data))
		case 4:
			var judul string
			fmt.Print("Masukkan Judul Game : ")
			fmt.Scanln(&judul)
			hasil := searchByJudul(data, judul)
			fmt.Println(hasil)
		case 5:
			top3 := data[0:3]
			fmt.Print(viewAllData(top3, "List Top 3 Game \n"))
		case 6:
			var rate float32
			fmt.Print("Masukkan minimal rating : ")
			fmt.Scanln(&rate)
			hasil := viewDataByRate(data, rate)
			fmt.Println(hasil)
		}
	}
}
