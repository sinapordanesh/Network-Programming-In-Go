package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gliderlabs/ssh"
)

func main() {
	ssh.Handle(handleCommands)
	log.Println("firing up the server...")
	log.Fatal(ssh.ListenAndServe("127.0.0.1:22", nil,
		ssh.HostKeyFile(filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")),
		ssh.PasswordAuth(
			ssh.PasswordHandler(func(ctx ssh.Context, password string) bool {
				return password == os.Getenv("SSH_PASSWORD")
			}),
		),
	))
}

func handleCommands(s ssh.Session) {
	switch s.RawCommand() {
		case "show version":
			ver := "Distributor ID: Ubuntu \n Description:    Ubuntu 22.04.2 LTS \n Release:        22.04 \n Codename:       jammy"
			io.WriteString(s, ver)
	}

}