package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Metode 1
	var input1 string

	fmt.Print("Masukkan Input Anda ")
	fmt.Scanln(&input1)

	// Metode 2 (Berakhiran \n)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("Hasil input ke-2 : ", text)

	ascii, _, _ := reader.ReadRune()
	fmt.Println("ASCII dari input", ascii)
}
