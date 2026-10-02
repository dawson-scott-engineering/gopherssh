package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

type Device struct {
	Host      string
	Username  string
	Password  string
	Platform  string
	SSHClient *ssh.Client
}

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

func (d *Device) Connect() error {
	config := d.setUpConfig()

	client, err := ssh.Dial("tcp", d.Host, &config)
	if err != nil {
		return err
	}

	d.SSHClient = client
	return nil
}

func (d *Device) RunCommand(command string) ([]byte, error) {
	session, err := d.SSHClient.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (d *Device) Close() {
	d.SSHClient.Close()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	username := os.Getenv("CISCO_USERNAME")
	password := os.Getenv("CISCO_PASSWORD")
	host := os.Getenv("CISCO_HOST")

	device := Device{
		Host:     host,
		Username: username,
		Password: password,
		Platform: "cisco_ios",
	}

	err = device.Connect()
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}

	output, err := device.RunCommand("show run")
	if err != nil {
		fmt.Println("Failed to run command:", err)
		device.Close()
		return
	}

	fmt.Println(string(output))

	device.Close()
}
