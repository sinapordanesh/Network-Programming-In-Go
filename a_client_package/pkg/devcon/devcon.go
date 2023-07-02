package devcon

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)

type sshClient struct {
	target string
	cfg    *ssh.ClientConfig
}

func NewClient(target string) *sshClient {
	return &sshClient{
		target: target,
		cfg: &ssh.ClientConfig{
			User: os.Getenv("SSH_USER"),
			Auth: []ssh.AuthMethod{
				ssh.Password(os.Getenv("SSH_PASSWORD")),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         time.Second * 5,
		},
	}
}

func (c *sshClient) run(target, cmd string) (string, error) {
	cfg := &ssh.ClientConfig{
		User: os.Getenv("SSH_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SSH_PASSWORD")),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:22", target), cfg)
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
