package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	SecretKey string `json:"secret_key"`
}

func readConfig() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error occured: %v\n", r)
			fmt.Println("application terminated gracefully....")
		} else {
			fmt.Println("application executed successfully....")
		}
	}()

	data, err := os.ReadFile("config.json")

	if err != nil {
		panic(fmt.Sprintf("error reading configuration file: %v\n", err))
	}

	var config Config
	err = json.Unmarshal(data, &config)

	if err != nil {
		panic(fmt.Sprintf("error parsing configuration file: %v\n", err))
	}

	if config.SecretKey == "" {
		panic("secret key is missing from configuration")
	}

	fmt.Println("configuration loaded successfully")
}

func main() {
	readConfig()
}
