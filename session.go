package gopherssh

import (
	"errors"
)

func (d *Device) RunCommand(command string) ([]byte, error) {
	if d.SSHClient == nil {
		return nil, errors.New("SSH client is not connected")
	}

	session, err := d.SSHClient.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return output, err
	}

	return output, nil
}
