package main

import (
	"fmt"
	"os"

	"github.com/nhereman/cloudflare-dns-updater/cloudflare"
	"github.com/nhereman/cloudflare-dns-updater/configuration"
	"github.com/nhereman/cloudflare-dns-updater/ip"
)

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("ERROR: failed to retrieve user home directory \n\t", err)
		os.Exit(1)
	}

	configFile := userHomeDir + "/.config/cloudflare-dns-updater/config.json"
	config, err := configuration.Load(configFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	err = configuration.Verify(config)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	publicIP, err := ip.Query()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	cloudflareAuth := cloudflare.CFAuth{
		Email:  config.Email,
		APIKey: config.CloudflareAPIKey,
	}

	record, err := cloudflare.GetRecord(cloudflareAuth, config.ZoneID, config.DNSRecordID)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	if record.Domain != config.DomainName {
		fmt.Println("ERROR:", "The domain configured ("+config.DomainName+") is not aligned with the one of the record ("+record.Domain+")")
		os.Exit(1)
	}

	if record.IP == publicIP {
		fmt.Println("INFO: public IP is already configured correctly. Stopping here.")
		os.Exit(0)
	}

	record.IP = publicIP

	err = cloudflare.SetRecord(cloudflareAuth, config.ZoneID, config.DNSRecordID, record)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}
