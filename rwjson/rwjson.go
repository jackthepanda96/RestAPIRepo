package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Alamat struct {
	Provinsi string `json:"provinsi"`
	Kota     string `json:"kota"`
}

type Pegawai struct {
	Nama   string   `json:"nama"`
	Alamat []Alamat `json:"alamat"`
	Email  string   `json:"email"`
	HP     int      `json:"hp"`
}

func writeJSON() {
	pegawai1 := Pegawai{"Jerry", []Alamat{{"Jawa Timur", "Surabaya"}, {"Jawa Tengah", "Semarang"}}, "jerry@alterra.id", 123456}

	file, _ := json.MarshalIndent(pegawai1, "", " ")

	err := ioutil.WriteFile("hasilcoba.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func readJSON() {
	tmpPegawai := Pegawai{}

	file, _ := ioutil.ReadFile("hasilcoba.json")
	err := json.Unmarshal(file, &tmpPegawai)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tmpPegawai)
}

func main() {
	fmt.Println("Writing ....")
	writeJSON()
	fmt.Println("Now Read")
	readJSON()
}
