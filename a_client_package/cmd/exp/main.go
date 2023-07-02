package main

import (
	"flag"
	"fmt"
	//"github.com/codered-by-ec-council/Hands-on-Network-Programming-with-Go/pkg/devcon"
	"a_clien_pkg/pkg/devcon"
	"log"
)

func main() {
	var target *string = flag.String("target", "127.0.0.1", "target against which to run a command")
	cmd := flag.String("cmd", "", "command to run against target device")
	flag.Parse()

	client := devcon.NewClient(*target)
	output, err := client.Run(*target, *cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
