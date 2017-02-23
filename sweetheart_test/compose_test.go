package sweetheart_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/options"
	"golang.org/x/net/context"
)

const DIR = "../"

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
		filename: path.Join(DIR, uniqId+".tmp.yml"),
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
