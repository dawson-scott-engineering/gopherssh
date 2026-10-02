package gopherssh

import (
	"errors"

	"golang.org/x/crypto/ssh"
)

func (d *Device) Connect() error {
	config := d.setUpConfig()

	client, err := ssh.Dial("tcp", d.Host, &config)
	if err != nil {
		return err
	}

	d.SSHClient = client
	return nil
}

func (d *Device) Close() error {
	if d.SSHClient == nil {
		return errors.New("SSH client is not connected")
	}

	return d.SSHClient.Close()
}
