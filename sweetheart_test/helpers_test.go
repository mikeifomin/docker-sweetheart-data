package sweetheart_test

import (
	"bytes"
	"html/template"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

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
