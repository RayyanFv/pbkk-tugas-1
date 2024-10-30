package main

import (
	"bufio"
	"fmt"
	"os"
)

// Fungsi untuk membalikkan string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Membaca input dari pengguna
	fmt.Print("Masukkan sebuah string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Menghasilkan string yang dibalik
	reversed := reverseString(input)
	fmt.Println("String setelah dibalik:", reversed)
}
