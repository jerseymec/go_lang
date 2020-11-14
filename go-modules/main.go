package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	fmt.Println("This is a module")
	log.Println("This is a log")
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
