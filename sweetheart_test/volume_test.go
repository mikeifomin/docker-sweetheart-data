package sweetheart_test

import "testing"

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
    build:
      context: .
      dockerfile: Dockerfile
    volumes:
      - data:/bkp-all/foo
      - data:/sync-only/foo
    environment:
      CRON: "@daily"
    
`)
  defer prj.Kill()
}
