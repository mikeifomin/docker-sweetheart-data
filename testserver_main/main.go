
package main

import (
	"github.com/mikeifomin/docker-sweetheart-data/testserver"
	"os"
	"path"
)

func main() {
	finish := make(chan bool)
	dir := os.Getenv("ROOT")
	if  dir == "" {
		ex, _ := os.Executable()
		dir = path.Dir(ex)
	}
	testserver.NewServer(3000, dir)
	<-finish
}
