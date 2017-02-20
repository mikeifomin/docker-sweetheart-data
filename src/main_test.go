package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/options"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"text/template"
	"time"
)

type compose struct {
	Port     int
	filename string
	prjName  string
	spec     string
	prj      *project.APIProject
}

func NewCompose(spec string) (*compose, error) {
	uniqId := fmt.Sprint(time.Now().UnixNano()) + randStr(5)
	log.Println(uniqId)
	c := compose{
		spec:     spec,
		Port:     findAvailablePort(3000),
		filename: uniqId + ".tmp.yml",
		prjName:  uniqId,
	}
	err := c.run()
	if err != nil {
		log.Println(err)
		c.clean()
		return nil, err
	}
	return &c, nil
}

func (c *compose) run() error {
	spec, errC := compile(c.spec, c)
	log.Println(spec)
	if errC != nil {
		return errC
	}
	_ = ioutil.WriteFile(c.filename, []byte(spec), 0644)
	prj, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeFiles: []string{c.filename},
			ProjectName:  c.prjName,
		},
	}, nil)

	if err != nil {
		return err
	}

	err = prj.Up(context.Background(), options.Up{})
	if err != nil {
		return err
	}
	c.prj = &prj
	return nil
}

func compile(str string, data interface{}) (string, error) {
	tmpl, err := template.New("my").Parse(str)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	errE := tmpl.Execute(&buf, data)
	if errE != nil {
		return "", errE
	}
	return buf.String(), nil
}

func (c *compose) clean() {
	_ = os.Remove(c.filename)
}
func (c *compose) Kill() {
	if c == nil {
		return
	}
	if c.prj != nil {
		(*c.prj).Down(context.Background(), options.Down{})
	}
	c.clean()
}

func findAvailablePort(port int) int {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fullHost := ln.Addr().String()
	parts := strings.Split(fullHost, ":")
	res, err := strconv.ParseInt(parts[len(parts)-1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(res)
}

func randStr(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func TestSelfDockerComposeRun(t *testing.T) {
	prj, _ := NewCompose(`
version: '2'
services:
  mongo:
    image: mongo:3.2
`)

	defer prj.Kill()
}

type AnswerHealth struct {
	Status string
}

func TestSelfTestserver(t *testing.T) {
	prj, _ := NewCompose(`
version: '2'
services:
  mongo:
    image: mongo:3.2
  testserver:
    build: ./testserver
    ports:
      - {{.Port}}:3000
`)
	time.Sleep(time.Second * 3)
	url := "http://localhost:" + strconv.Itoa(prj.Port) + "/health"
	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
		return
	}
	defer res.Body.Close()

	var answer AnswerHealth
	_ = json.NewDecoder(res.Body).Decode(&answer)
	if answer.Status != "ok" {
		t.Error("test server does not responding")
	}
	defer prj.Kill()
}
