package gopherssh

import (
	"strings"
	"testing"
)

func TestRunCommandNotConnected(t *testing.T) {
	device := Device{}

	_, err := device.RunCommand("show version")

	if err == nil {
		t.Fatal("expected error when device is not connected")
	}

	if !strings.Contains(err.Error(), "not connected") {
		t.Errorf("expected error to contain %q, got %q", "not connected", err)
	}
}

func TestRunCommand(t *testing.T) {
	device := newTestDevice(t)

	err := device.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer device.Close()

	output, err := device.RunCommand("show version")

	if err != nil {
		t.Fatal(err)
	}

	if len(output) == 0 {
		t.Fatal("expected command output")
	}

	if !strings.Contains(string(output), "Cisco IOS") {
		t.Errorf("expected output to contain %q, got:\n%s", "Cisco IOS", output)
	}
}

func TestRunCommandFailure(t *testing.T) {
	device := newTestDevice(t)

	err := device.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer device.Close()

	output, err := device.RunCommand("this-is-not-a-real-command")

	if err != nil {
		t.Errorf("unexpected SSH error: %v", err)
	}

	if len(output) == 0 {
		t.Fatal("expected output from failed command")
	}

	if !strings.Contains(string(output), "invalid autocommand") {
		t.Errorf("expected Cisco CLI error output, got:\n%s", output)
	}
}