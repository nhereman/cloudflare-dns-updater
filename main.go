package main

import (
	"fmt"
	"os"

	"github.com/nhereman/cloudflare-dns-updater/cloudflare"
	"github.com/nhereman/cloudflare-dns-updater/configuration"
	"github.com/nhereman/cloudflare-dns-updater/ip"
)

func main() {
	configurationFile := "/home/nhereman/.config/cloudflare-dns-updater/config.json"
	configuration, err := configuration.Load(configurationFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	publicIP, err := ip.Query()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	cloudflareAuth := cloudflare.CFAuth{
		Email:  configuration.Email,
		APIKey: configuration.CloudflareAPIKey,
	}

	record, err := cloudflare.GetRecord(cloudflareAuth, configuration.ZoneID, configuration.DNSRecordID)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	if record.Domain != configuration.DomainName {
		fmt.Println("ERROR:", "The domain configured ("+configuration.DomainName+") is not aligned with the one of the record ("+record.Domain+")")
		os.Exit(1)
	}

	if record.IP == publicIP {
		fmt.Println("INFO: public IP is already configured correctly. Stopping here.")
		os.Exit(0)
	}

	record.IP = publicIP

	err = cloudflare.SetRecord(cloudflareAuth, configuration.ZoneID, configuration.DNSRecordID, record)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}
