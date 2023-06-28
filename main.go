package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.New()
	log.Print("hello world")
}