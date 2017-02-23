package sweetheart_test

import (
	"testing"
	"time"
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
    build:
      context: .
      dockerfile: Dockerfile.testserver
    ports:
      - {{.Port}}:3000
`)
	time.Sleep(time.Second * 3)
	defer prj.Kill()
}
func TestDataSync(t *testing.T) {

	prj, _ := NewCompose(`
version: '2'
volumes:
  data:
services:
  mongo:
    image: mongo:3.2
  testserver:
    build:
      context: .
      dockerfile: Dockerfile.testserver
    volumes:
     - data:/data
    ports:
      - {{.Port}}:3000
  bkp:
    build: .
    volumes:
      - data:/bkp-all/foo
      - data:/sync-only/foo
    environment:
      CRON: "@daily"
    
`)
	defer prj.Kill()
}
