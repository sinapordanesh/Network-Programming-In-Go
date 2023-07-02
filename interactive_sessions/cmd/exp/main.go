package main

import (
	"flag"
	"fmt"
	"interactive_sessions/pkg/devcon"
	"log"
)

func main() {
	target := flag.String("target", "127.0.0.1", "target against which to run a command")
	// cmd := flag.String("cmd", "", "command to run against target device")
	flag.Parse()
	client := devcon.NewClient(*target)
	cmds := []string{
		"configure",
		"set system host-name MODIFIED_HOSTNAME",
		"commit and-quit comment MODIFIED HOSTNAME",
		"exit",
	}
	output, err := client.RunAll(cmds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
