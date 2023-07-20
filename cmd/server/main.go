package main

import (
	"kiplingkelvin/golang-skeleton/internal/server"

	"github.com/sirupsen/logrus"
)

func main() {

	err := server.RunServer()
	if err != nil {
		logrus.WithField("Error", err).Fatal(err)
	}

}
