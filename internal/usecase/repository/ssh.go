package usecase_repository

type SSHClient interface {
	Execute(cmd string) (string, error)
	Close() error
}

type SSHClientFactory interface {
	NewSSHClient(user string, addr string, password string) (SSHClient, error)
}