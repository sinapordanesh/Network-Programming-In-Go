package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"log"

	"golang.org/x/crypto/ssh"
	"github.com/pkg/errors"
)

func main() {
	target := flag.String("target", "127.0.0.1", "Target to connect to run a command")
	cmd := flag.String("cmd", "", "Command to run against the target device")
	flag.Parse()
	output, err := run(*target, *cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

}

func run(target, cmd string) (string, error) {
	cfg := &ssh.ClientConfig{
		User: os.Getenv("SSH_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSH_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), 
		Timeout: time.Second*5,
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:2222",target), cfg)
	if err != nil {
		return "", errors.Wrap(err, "dial failed")
	
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(output), nil

}
