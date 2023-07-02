package devcon

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"time"
)

type sshClient struct {
	target string
	port   string
	cfg    *ssh.ClientConfig
}

type option func(*sshClient)

func SetPort(port string) option {
	return func(c *sshClient) {
		c.port = port
	}
}

func SetPassword(pw string) option {
	return func(c *sshClient) {
		authMethod := []ssh.AuthMethod{
			ssh.Password(pw),
		}
		c.cfg.Auth = authMethod
	}
}

func NewClient(user, target string, opts ...option) *sshClient {
	defaultPort := "22"
	client := &sshClient{
		port:   defaultPort,
		target: target,
		cfg: &ssh.ClientConfig{
			User:            user,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         time.Second * 5,
		},
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *sshClient) Run(target, cmd string) (string, error) {
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

func (c *sshClient) RunAll(cmds []string) (string, error) {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:22", c.target), c.cfg)
	if err != nil {
		return "", errors.Wrap(err, "dial failed")
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	wc, err := session.StdinPipe()
	if err != nil {
		return "", err
	}
	defer wc.Close()
	r, err := session.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := session.Shell(); err != nil {
		return "", err
	}
	for _, cmd := range cmds {
		_, err := fmt.Fprintf(wc, "%s\\n", cmd)
		if err != nil {
			return "", err
		}
	}
	if err := session.Wait(); err != nil {
		return "", err
	}
	bs, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
