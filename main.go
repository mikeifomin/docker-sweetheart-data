package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	neverEnds := make(chan bool)
	fmt.Println("hi")
	c := cron.New()
	err := c.AddFunc("* * * * * *", func() { fmt.Println("yay") })
	if err != nil {
		panic(err)
	}

	c.Start()
	<-neverEnds
}
