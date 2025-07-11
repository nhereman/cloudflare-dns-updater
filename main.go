package main

import (
	"fmt"

	"github.com/nhereman/cloudflare-dns-updater/ip"
)

func main() {
	publicIP, err := ip.Query()
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println(publicIP)
}
