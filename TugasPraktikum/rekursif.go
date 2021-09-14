package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func DeretRekursifBilanganReal(i float32, n float32, max float32) float32 {
	var total float32 = 0.0
	if n < max {
		bagi := i / n
		total += bagi
		fmt.Printf("%.2f + ", bagi)
		return bagi + DeretRekursifBilanganReal(i, n+1, max)
	} else {
		bagi := i / n
		total += bagi
		fmt.Printf("%.2f ", bagi)
		return bagi
	}
}

func deretBerpangkat(n int, max int, posisi string) string {
	if n < max {
		pangkat := int(math.Pow(float64(n), float64(2)))
		if posisi == "atas" {
			fmt.Printf("%dx/%d + ", pangkat, n)
		} else {
			fmt.Printf("%d/%dx + ", n, pangkat)
		}
		return deretBerpangkat(n+1, max, posisi)
	} else {
		pangkat := int(math.Pow(float64(n), float64(2)))
		if posisi == "atas" {
			fmt.Printf("%dx/%d ", pangkat, n)
		} else {
			fmt.Printf("%d/%dx ", n, pangkat)
		}
		return ""
	}
}

func countNonCapz(word string, n int) int {
	low := strings.ToLower(word)
	if n < len(word) {
		var same bool = false
		if string(word[n]) == string(low[n]) {
			same = true
		}
		if same {
			return 1 + countNonCapz(word, n+1)
		} else {
			return 0 + countNonCapz(word, n+1)
		}
	} else {
		return 0
	}
}

func countNumberInString(word string, n int) int {
	if n < len(word) {
		_, err := strconv.Atoi(string(word[n]))
		if err == nil {
			return 1 + countNumberInString(word, n+1)
		} else {
			return 0 + countNumberInString(word, n+1)
		}
	} else {
		return 0
	}
}

func main() {
	fmt.Print("a. SUM = ")
	fmt.Println("=", DeretRekursifBilanganReal(20, 1, 5))
	fmt.Print("b. SUM = ")
	fmt.Println("=", DeretRekursifBilanganReal(100, 1, 5))
	fmt.Print("c. SUM = ")
	fmt.Println("=", DeretRekursifBilanganReal(1, 1, 5))
	fmt.Print("d. SUM = ")
	fmt.Println("=", deretBerpangkat(1, 5, "bawah"))
	fmt.Print("e. SUM = ")
	fmt.Println("=", deretBerpangkat(1, 5, "atas"))

	var word string = "Makan"
	fmt.Printf("\nCount non Capitalize '%s' = ", word)
	fmt.Println(countNonCapz(word, 0))

	var newWord string = "JKT48"
	fmt.Printf("\nCount number string '%s' = ", newWord)
	fmt.Println(countNumberInString(newWord, 0))
}
