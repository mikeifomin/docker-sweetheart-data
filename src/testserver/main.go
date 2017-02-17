package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	DIR = "/data"
)

type answer struct {
	Files []string
}

func main() {
	http.HandleFunc("/show_files", func(w http.ResponseWriter, r *http.Request) {
		files, err := ioutil.ReadDir(DIR)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		answ := answer{make([]string, 1)}
		for _, file := range files {
			fmt.Println(file)
			answ.Files = append(answ.Files, file.Name())
		}
		jsonData, _ := json.Marshal(answ)
		w.Write(jsonData)
	})
	http.HandleFunc("/add_file", func(w http.ResponseWriter, r *http.Request) {
		filename := "1.json"
		data := []byte("{}")
		ioutil.WriteFile(path.Join(DIR, filename), data, 0777)
	})
	http.ListenAndServe(":3000", nil)
}
