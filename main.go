package main

import (
	"github.com/go-rendezvous/rendezvous/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	if err = cmd.RootCommand.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
