package main

import (
	"./testserver"
	"os"
	"path"
	"path/filepath"
)

func main() {
	finish := make(chan bool)
	if dir := os.Getenv("ROOT"); dir == "" {
		ex, _ := os.Executable()
		dir = path.Dir(ex)
	}
	testserver.NewServer(3000, dir)
	<-finish
}
