package testserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
)

type Srv struct {
	Srv  http.ServeMux
	Port int
	Root string
}

func NewServer(port int, root string) *Srv {
	srv := http.NewServeMux()
	srv.HandleFunc("/health", health)
	srv.HandleFunc("/addFile", addFile)
	srv.HandleFunc("/listFiles", listFiles)
	bind := ":" + strconv.Itoa(port)
	go func() {
		_ = http.ListenAndServe(bind, srv)
	}()
	result := Srv{*srv, port, root}
	return &result
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(RespHealth{"ok"})
}

func listFiles(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(DIR)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	result := RespListFile{}
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
