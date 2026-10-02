package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dawson-scott-engineering/gopherssh"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env:", err)
	}

	device := gopherssh.Device{
		Host:     os.Getenv("CISCO_HOST"),
		Username: os.Getenv("CISCO_USERNAME"),
		Password: os.Getenv("CISCO_PASSWORD"),
		Platform: "cisco_ios",
	}

	if err := device.Connect(); err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer device.Close()

	output, err := device.RunCommand("show run")
	if err != nil {
		log.Println("Failed to run command:", err)
	}

	fmt.Println(string(output))
}
