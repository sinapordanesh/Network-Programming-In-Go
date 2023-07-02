package main

import (
	"clien_options/pkg/devcon"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	target := flag.String("target", "127.0.0.1", "target against which to run a command")
	cmdFile := flag.String("cmdfile", "", "command filename")
	flag.Parse()

	client := devcon.NewClient(*target, devcon.SetPassword(os.Getenv("SSH_PASSWORD")))

	f, err := os.Open(*cmdFile)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	cmds := strings.Split(string(bs), "\\n")
	output, err := client.RunAll(cmds)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
