package repository

import (
	"bytes"
	"net"

	"github.com/Goalt/hostmonitor/internal/usecase/repository"
	"golang.org/x/crypto/ssh"
)

type SSHFactory struct {
}

func NewSSHFactory() *SSHFactory {
	return &SSHFactory{}
}

func (s *SSHFactory) NewSSHClient(user string, addr string, password string) (usecase_repository.SSHClient, error) {
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	// Connect
	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, "22"), config)
	if err != nil {
		return nil, err
	}

	return &SSHClient{client}, nil
}

type SSHClient struct {
	*ssh.Client
}

func (c *SSHClient) Execute(cmd string) (string, error) {
	// Create a session. It is one session per command.
	session, err := c.Client.NewSession()
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	session.Stdout = &b

	err = session.Run(cmd)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (c *SSHClient) Close() error {
	return c.Client.Close()
}
