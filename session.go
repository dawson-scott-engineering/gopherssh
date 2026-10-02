package gopherssh

import (
	"fmt"
)

func (d *Device) RunCommand(command string) ([]byte, error) {

	if d.SSHClient == nil {
		return nil, fmt.Errorf("the client for %s is not connected", d.Host)
	}

	session, err := d.SSHClient.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session for %s: %w", d.Host, err)
	}

	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return output, fmt.Errorf("failed to run command on %s: %w", d.Host, err)
	}

	return output, nil
}
