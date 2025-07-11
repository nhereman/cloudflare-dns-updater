// Package configuration handle the configuration of the cloudflare-dns-updater tool
package configuration

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	domainNameEnvVar       = "CDU_DOMAIN_NAME"
	dnsRecordIDEnvVar      = "CDU_DNS_RECORD_ID"
	emailEnvVar            = "CDU_CF_EMAIL"
	cloudflareAPIKeyEnvVar = "CDU_CF_API_KEY"
	zoneIDEnvVar           = "CDU_CF_ZONE_ID"
)

type Configuration struct {
	DomainName       string `json:"domain"`
	DNSRecordID      string `json:"record"`
	Email            string `json:"email"`
	CloudflareAPIKey string `json:"api_key"`
	ZoneID           string `json:"zone"`
}

func Load(configurationFile string) (Configuration, error) {
	configuration, err := loadFromFile(configurationFile)
	if err != nil {
		return Configuration{}, errors.New("Failed to get configuration from file " + configurationFile + ": \n\t" + err.Error())
	}

	domain := os.Getenv(domainNameEnvVar)
	if domain != "" {
		configuration.DomainName = domain
	}

	recordID := os.Getenv(dnsRecordIDEnvVar)
	if recordID != "" {
		configuration.DNSRecordID = recordID
	}

	email := os.Getenv(emailEnvVar)
	if email != "" {
		configuration.Email = email
	}

	apiKey := os.Getenv(cloudflareAPIKeyEnvVar)
	if apiKey != "" {
		configuration.CloudflareAPIKey = apiKey
	}

	zone := os.Getenv(zoneIDEnvVar)
	if zone != "" {
		configuration.ZoneID = zone
	}

	return configuration, nil
}

func Verify(configuration Configuration) error {
	if configuration.DomainName == "" {
		return errors.New("domain configuration missing")
	}
	if configuration.DNSRecordID == "" {
		return errors.New("record configuration missing")
	}
	if configuration.Email == "" {
		return errors.New("email configuration missing")
	}
	if configuration.CloudflareAPIKey == "" {
		return errors.New("api_key configuration missing")
	}
	if configuration.ZoneID == "" {
		return errors.New("zone configuration missing")
	}
	return nil
}

func loadFromFile(configurationFile string) (Configuration, error) {
	if configurationFile == "" {
		return Configuration{}, nil
	}

	content, err := os.ReadFile(configurationFile)
	if err != nil {
		return Configuration{}, errors.New("Failed to open " + configurationFile + ": \n\t" + err.Error())
	}

	var configuration Configuration
	err = json.Unmarshal(content, &configuration)
	if err != nil {
		return Configuration{}, errors.New("Failed to decode json: \n\t" + err.Error())
	}

	return configuration, nil
}
