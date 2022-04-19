package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Films     []string `json:"films"`
}

func getStatic(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "text/json")

	switch req.Method {
	case "GET":
		fmt.Println(req.URL.Query().Get("category"))
		pegawai1 := Pegawai{"Jerry", []Alamat{{"Jawa Timur", "Surabaya"}, {"Jawa Tengah", "Semarang"}}, "jerry@alterra.id", 123456}
		// data := "Ini Method GET"
		file, _ := json.Marshal(pegawai1)

		w.Write(file)
	case "POST":
		data := "Ini Method POST"
		file, _ := json.Marshal(data)
		tmp := Pegawai{}
		readBody := json.NewDecoder(req.Body)
		readBody.Decode(&tmp)
		fmt.Println(tmp)

		w.Write(file)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		data := "Method ini tidak di dukung"
		file, _ := json.Marshal(data)
		w.Write(file)
	}

	return
}

func readSwapi() {
	res, err := http.Get("https://swapi.dev/api/people/2")
	if err != nil {
		fmt.Println(err)
	}

	tmp := StarWarsPeople{}

	err = json.NewDecoder(res.Body).Decode(&tmp)
	fmt.Println(tmp)
}

func main() {
	readSwapi()

	http.HandleFunc("/", getStatic)
	http.HandleFunc("/halo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			for key, val := range r.URL.Query() {
				fmt.Println(key, val)
			}
			fmt.Println(r.URL.Query())
			fmt.Println(r.URL.Query().Get("page"))
		}
	})
	http.HandleFunc("/halo/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte("Ini /halo/"))
		}
	})

	fmt.Println(http.ListenAndServe(":8000", nil))

}
