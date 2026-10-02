package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	connectStart := time.Now()

	if err := device.Connect(); err != nil {
		log.Fatal("Failed to connect:", err)
	}

	fmt.Println("Connect took:", time.Since(connectStart))

	defer device.Close()

	commandStart := time.Now()

	output, err := device.RunCommand("sh ip int br")
	if err != nil {
		log.Println("Failed to run command:", err)
	}

	fmt.Println("Command took:", time.Since(commandStart))
	fmt.Println(string(output))
}
