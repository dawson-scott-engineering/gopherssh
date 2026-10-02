package gopherssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// Connect establishes an SSH connection to the device.

func (d *Device) Connect() error {

	if d.SSHClient != nil {
		return fmt.Errorf("%s already has a connection", d.Host)
	}

	config := d.setUpConfig()

	client, err := ssh.Dial("tcp", d.Host, &config)
	if err != nil {
		return fmt.Errorf("connecting to %s: %w", d.Host, err)
	}

	d.SSHClient = client
	return nil
}

// Close closes the device's SSH connection and clears the client.
// If the device is not connected, Close returns nil.

func (d *Device) Close() error {

	if d.SSHClient == nil {
		return nil
	}

	err := d.SSHClient.Close()

	if err != nil {
		d.SSHClient = nil
		return fmt.Errorf("error closing connection to %s: %w", d.Host, err)
	}
	d.SSHClient = nil
	return nil

}
