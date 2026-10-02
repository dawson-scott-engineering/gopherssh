package gopherssh

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("cmd/ssh-test/.env"); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func newTestDevice(t *testing.T) Device {
	t.Helper()

	return Device{
		Host:     os.Getenv("CISCO_HOST"),
		Username: os.Getenv("CISCO_USERNAME"),
		Password: os.Getenv("CISCO_PASSWORD"),
		Platform: "cisco_ios",
	}
}

func TestConnect(t *testing.T) {
	device := newTestDevice(t)

	err := device.Connect()
	if err != nil {
		t.Fatal(err)
	}

	defer device.Close()

	if device.SSHClient == nil {
		t.Error("expected SSHClient to be set after connecting")
	}
}

func TestConnectAlreadyConnected(t *testing.T) {
	device := newTestDevice(t)

	err := device.Connect()
	if err != nil {
		t.Fatal(err)
	}

	defer device.Close()

	err = device.Connect()

	if err == nil {
		t.Error("expected an error when connecting twice")
	}
}

func TestConnectFailure(t *testing.T) {
	device := newTestDevice(t)

	device.Password = "definitely-wrong-password"

	err := device.Connect()

	if err == nil {
		t.Error("expected connection error")
	}

	if device.SSHClient != nil {
		t.Error("expected SSHClient to remain nil after failed connection")
	}
}

func TestCloseNotConnected(t *testing.T) {
	device := Device{}

	err := device.Close()

	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestClose(t *testing.T) {
	device := newTestDevice(t)

	err := device.Connect()
	if err != nil {
		t.Fatal(err)
	}

	err = device.Close()
	if err != nil {
		t.Fatal(err)
	}

	if device.SSHClient != nil {
		t.Error("expected SSHClient to be nil after closing")
	}
}
