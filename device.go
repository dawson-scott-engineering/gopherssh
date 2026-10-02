package gopherssh

import "golang.org/x/crypto/ssh"

type Device struct {
	Host      string
	Username  string
	Password  string
	Platform  string
	SSHClient *ssh.Client
}
