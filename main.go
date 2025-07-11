package main

import (
	"fmt"

	"github.com/nhereman/cloudflare-dns-updater/configuration"
	"github.com/nhereman/cloudflare-dns-updater/ip"
)

func main() {
	configurationFile := "/home/nhereman/.config/cloudflare-dns-updater/config.json"
	configuration, err := configuration.Load(configurationFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("Configuration:", configuration)

	publicIP, err := ip.Query()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("ip: ", publicIP)
}
