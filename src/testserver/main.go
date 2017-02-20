package main

import (
	"encoding/json"
	"fmt"
	"log"
	//"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	DIR = "/data"
)

type OptNewFile struct {
	Filename  string
	Contents  string
	Overwrite bool
}

type File struct {
	Filename string
	Contents string
}
type AnswerListFiles struct {
	Files []File
}
type AnswerHealth struct {
	Status string
}
type AnswerAddFile struct {
	IsOk  bool
	Write int
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(AnswerHealth{"ok"})
}

func listFiles(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	result := AnswerListFiles{}
	for _, f := range files {
		fmt.Println(f)
		file := File{f.Name(), ""}
		result.Files = append(result.Files, file)
	}
	json.NewEncoder(w).Encode(result)
}

func addFile(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var opt OptNewFile
	err := decoder.Decode(&opt)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	data := []byte(opt.Contents)
	errW := ioutil.WriteFile(path.Join(DIR, opt.Filename), data, 0777)
	if errW != nil {
		w.WriteHeader(500)
	}

}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/addFile", addFile)
	http.HandleFunc("/listFiles", listFiles)
	http.ListenAndServe(":3000", nil)
}
