package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var (
	config     Config
	configBody string
	configPath string
)

func init() {
	flag.StringVar(&configBody, "config", "", "Config json for application")
	flag.StringVar(&configPath, "config-path", "", "Path to config json for application")

	flag.Parse()
}

func main() {
	if configBody == "" {
		if configPath == "" {
			printErrorAndExit(ErrorMissingConfig)
		}
		buff, err := os.ReadFile(configPath)
		printErrorAndExit(err)
		configBody = string(buff)
	}
	err := json.Unmarshal([]byte(configBody), &config)
	if err != nil {
		printErrorAndExit(ErrorBadConfig)
	}
}

func printErrorAndExit(err error) {
	if err == nil {
		return
	}
	log.Println(err)
	os.Exit(1)
}
