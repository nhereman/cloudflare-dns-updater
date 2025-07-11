package main

import (
	"flag"
	"log"
	"os"

	"github.com/nhereman/cloudflare-dns-updater/cloudflare"
	"github.com/nhereman/cloudflare-dns-updater/configuration"
	"github.com/nhereman/cloudflare-dns-updater/ip"
)

func main() {
	log.Print("INFO: executing cloudflare-dns-updater")

	var configFile string
	flag.StringVar(&configFile, "c", "", "Specify configuration file path. Default: ~/.config/cloudflare-dns-updater/config.json")
	flag.Parse()

	if configFile == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("ERROR: failed to retrieve user home directory", err)
		}
		configFile = userHomeDir + "/.config/cloudflare-dns-updater/config.json"
	}

	log.Print("INFO: configuration file:", configFile)

	config, err := configuration.Load(configFile)
	if err != nil {
		log.Fatal("ERROR: failed to load configuration", err)
	}

	err = configuration.Verify(config)
	if err != nil {
		log.Fatal("ERROR: configuration is incorrect", err)
	}

	publicIP, err := ip.Query()
	if err != nil {
		log.Fatal("ERROR: failed to get public IP", err)
	}

	cloudflareAuth := cloudflare.CFAuth{
		Email:  config.Email,
		APIKey: config.CloudflareAPIKey,
	}

	record, err := cloudflare.GetRecord(cloudflareAuth, config.ZoneID, config.DNSRecordID)
	if err != nil {
		log.Fatal("ERROR: failed to get DNS record from cloudflare", err)
	}

	if record.Domain != config.DomainName {
		log.Fatal("ERROR: the domain configured (" + config.DomainName + ") is not aligned with the one of the record (" + record.Domain + ")")
	}

	if record.IP == publicIP {
		log.Print("INFO: public IP is already configured correctly. Stopping here.")
		os.Exit(0)
	}

	record.IP = publicIP

	err = cloudflare.SetRecord(cloudflareAuth, config.ZoneID, config.DNSRecordID, record)
	if err != nil {
		log.Fatal("ERROR: failed to request a modification of the DNS record")
	}

	log.Print("INFO: DNS Record Updated with following IP: " + record.IP)
}
