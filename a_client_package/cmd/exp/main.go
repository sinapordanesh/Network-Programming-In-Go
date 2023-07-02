package main

import (
	"flag"
	"fmt"
	"github.com/codered-by-ec-council/Hands-on-Network-Programming-with-Go/pkg/devcon"
	"log"
)

func main() {
	var target *string = flag.String("target", "127.0.0.1", "target against which to run a command")
	cmd := flag.String("cmd", "", "command to run against target device")
	flag.Parse()

	client, _ := devcon.NewClient("go", *target)
	output, err := client.Run(*cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
