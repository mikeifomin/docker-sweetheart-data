package sweetheart_test

import (
	"fmt"
	"testing"
	"time"
	"github.com/mikeifomin/docker-sweetheart-data/testserver"
)

func TestSelfDockerComposeRun(t *testing.T) {
	prj, _ := NewCompose(`
version: '2'
services:
  mongo:
    image: mongo:3.2
`)

	defer prj.Kill()
}

func TestSelfTestserver(t *testing.T) {
	prj, _ := NewCompose(`
version: '2'
services:
  mongo:
    image: mongo:3.2
  testserver:
    build:
      context: .
      dockerfile: Dockerfile.testserver
    ports:
      - {{.Port}}:3000
`)
  resp := testserver.RespHealth{}
	err := prj.CallTestServer("/health",testserver.ParamsHealth{},&resp)
	if err != nil {
		t.Error(err)
	}
	if resp.Status != "ok" {
		t.Error("status not ok")
	}
	fmt.Println(prj)
	time.Sleep(time.Second * 3)
	fmt.Println(prj)
	defer prj.Kill()
}
