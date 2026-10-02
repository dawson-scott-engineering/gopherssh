package gopherssh

import "golang.org/x/crypto/ssh"

// Gens config for SSH Client

func (d *Device) setUpConfig() ssh.ClientConfig {
	sshConfig := ssh.ClientConfig{
		User: d.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(d.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return sshConfig
}
