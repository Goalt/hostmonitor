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
        User: user,
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

    // Create a session. It is one session per command.
    session, err := client.NewSession()
    if err != nil {
        return nil, err
    }

	return &SSHClient{session}, nil
}

type SSHClient struct {
	*ssh.Session
}

func (c *SSHClient) Execute(cmd string) (string, error) {
	var b bytes.Buffer 
	c.Stdout = &b

	err := c.Run(cmd)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (c *SSHClient) Close() error {
	return c.Session.Close()
}