package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile(text string) {
	myData := []byte(text)
	err := ioutil.WriteFile("filesaya.txt", myData, 0777)

	if err != nil {
		fmt.Println(err)
	}
}

func readFile() {
	data, err := ioutil.ReadFile("filesaya.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}

func main() {
	fmt.Println("Masukkan Input")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println("Hasil input adalah : ", text)
	fmt.Println("Write File")
	writeFile(text)
	fmt.Println("Read File")
	readFile()
}
