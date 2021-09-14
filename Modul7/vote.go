package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sort"
)

type music struct {
	judul    string
	penyanyi string
	vote     int
	id       string
}

func randToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func cekID(data map[string]music, id string) (string, bool) {
	_, notExist := data[id]
	return id, notExist
}

func required(input string) (string, bool) {
	var ok bool = input != ""
	return input, ok
}

func addData(data map[string]music, judul string, penyanyi string, vote int) map[string]music {
	var temp music
	var ok bool
	temp.judul = judul
	id, _ := randToken(3)
	id, ok = cekID(data, id)
	// looping sampe dapat id yang berbeda
	for ok {
		id, _ = randToken(3)
		_, check := cekID(data, id)
		ok = check
	}
	temp.id = id
	// temp.id, _ = randToken(3)
	temp.penyanyi = penyanyi
	temp.vote = vote

	data[temp.id] = temp

	return data
}

func deleteData(data map[string]music, id string) (map[string]music, bool) {
	_, ok := data[id]
	var hasil bool
	if ok {
		delete(data, id)
		hasil = true
	} else {
		hasil = false
	}
	return data, hasil
}

func viewAllData(data map[string]music, title string) string {
	var arrData []music
	for _, v := range data {
		arrData = append(arrData, v)
	}

	sort.Slice(arrData, func(i, j int) bool {
		return arrData[i].vote > arrData[j].vote
	})

	var hasil string = title
	for i := 0; i < len(arrData); i++ {
		hasil += fmt.Sprintf("%d. ID %s, %s oleh %s mendapatkan %d Vote \n", i+1, arrData[i].id, arrData[i].judul, arrData[i].penyanyi, arrData[i].vote)
	}

	return hasil

}

func hitungTotalVoting(data map[string]music) int {
	var total int = 0
	for _, v := range data {
		total += v.vote
	}
	return total
}

func viewTop3(data map[string]music, title string) string {
	var arrData []music
	for _, v := range data {
		arrData = append(arrData, v)
	}

	sort.Slice(arrData, func(i, j int) bool {
		return arrData[i].vote > arrData[j].vote
	})

	var hasil string = title
	many := len(data)
	if many > 3 {
		many = 3
	}
	for i := 0; i < many; i++ {
		hasil += fmt.Sprintf("%d. ID %s, %s oleh %s mendapatkan %d Vote \n", i+1, arrData[i].id, arrData[i].judul, arrData[i].penyanyi, arrData[i].vote)
	}

	return hasil

}

func findMusicBySinger(data map[string]music, word string) string {
	rexString := fmt.Sprintf("(?i)%s.||(?i)%s", word, word)
	regex := regexp.MustCompile(rexString)
	var hasil string = ""
	var i int = 0
	for _, v := range data {
		match := regex.Match([]byte(v.penyanyi))
		if match {
			hasil += fmt.Sprintf("%d. ID %s, %s oleh %s mendapatkan %d Vote \n", i+1, v.id, v.judul, v.penyanyi, v.vote)
			i++
		}
	}
	return hasil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// in := bufio.NewReader(os.Stdin)

	var pilihan int = 1
	var data map[string]music
	data = map[string]music{}
	for pilihan != 0 {
		fmt.Println(`Menu
1. Input Data Vote Musik
2. Hapus Data Vote Musik berdasarkan ID
3. Menampilkan Seluruh Data dengan Jumlah Data Yang Tersimpan
4. Total voting
5. Top 3 Musik Favorite
6. Cari Musik Berdasarkan Nama Penyanyi
0. Exit`)
		fmt.Print("Silahkan Pilih Menu : ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			var judul, penyanyi string
			var vote int
			var ok bool
			fmt.Print("Masukkan Judul Music : ")
			scanner.Scan()
			judul = scanner.Text()
			judul, ok = required(judul)
			if !ok {
				fmt.Println("Harap masukkan Judul, silahkan coba lagi")
			} else {
				fmt.Print("Masukkan Nama Penyanyi : ")
				scanner.Scan()
				penyanyi = scanner.Text()
				penyanyi, ok = required(penyanyi)
				if !ok {
					fmt.Println("Harap masukkan Penyanyi, silahkan coba lagi")
				} else {
					fmt.Print("Banyaknya Vote : ")
					fmt.Scanln(&vote)
					data = addData(data, judul, penyanyi, vote)
					fmt.Printf("Data berhasil ditambahkan \n \n")
				}
			}
		case 2:
			fmt.Println(viewAllData(data, "List Music \n"))
			var idx string
			fmt.Print("ID data yang ingin dihapus : ")
			fmt.Scanln(&idx)
			newData, status := deleteData(data, idx)
			if status {
				fmt.Printf("Data dengan id %s berhasil dihapus \n\n", idx)
				data = newData
			} else {
				fmt.Printf("Data dengan id %s tidak terdaftar \n\n", idx)
			}
			fmt.Print(viewAllData(data, "List Music \n"), "\n\n")
		case 3:
			fmt.Print(viewAllData(data, "List Music \n"))
			fmt.Printf("Total Data : %d \n\n", len(data))
		case 4:
			hasil := hitungTotalVoting(data)
			fmt.Printf("Total Voting : %d \n\n", hasil)
		case 5:
			fmt.Print(viewTop3(data, "List Top 3 Musik \n"), "\n\n")
		case 6:
			var word string
			fmt.Print("Masukkan huruf : ")
			fmt.Scanln(&word)
			hasil := findMusicBySinger(data, word)
			fmt.Println(hasil)
		}
	}
}
