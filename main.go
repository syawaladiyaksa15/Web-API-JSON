package main

import (
	// "fmt"
	"encoding/json"
	"net/http"
)

type mahasiswa struct {
	NPM   string
	Nama  string
	Grade int
}

var data = []mahasiswa{
	mahasiswa{"2015030031", "Syawal Adiyaksa", 23},
	mahasiswa{"2015030032", "Ryan Febrian", 22},
	mahasiswa{"2015010043", "Riska Malik", 23},
	mahasiswa{"2015020013", "Andi Lukman", 23},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var npm = r.FormValue("npm")
		var result []byte
		var err error

		for _, each := range data {
			if each.NPM == npm {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User tidak ditemukan", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	http.ListenAndServe(":8001", nil)
}
