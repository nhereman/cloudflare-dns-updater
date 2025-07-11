// Package cloudflare provides utilities to contact cloudflare API
package cloudflare

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type RecordDetailResponse struct {
	Error  []ErrorResponse `json:"errors"`
	Result RecordDetail    `json:"result"`
}

type RecordDetail struct {
	Domain   string `json:"name"`
	IP       string `json:"content"`
	TTL      int    `json:"ttl"`
	Type     string `json:"type"`
	Comment  string `json:"comment"`
	Proxied  bool   `json:"proxied"`
	Settings struct {
		IPV4Only bool `json:"ipv4_only"`
		IPV6Only bool `json:"ipv6_only"`
	} `json:"settings"`
	Tags      []string `json:"tags"`
	Proxiable bool     `json:"proxiable"`
}

func GetRecord(auth CFAuth, zoneID string, recordID string) (RecordDetail, error) {
	client := &http.Client{}
	url := apiURL + "/" + zoneID + "/" + dnsRecordsPath + "/" + recordID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RecordDetail{}, errors.New("failed to build request: \n\t" + err.Error())
	}
	req.Header.Set(emailHeader, auth.Email)
	req.Header.Set(apiKeyHeader, auth.APIKey)
	res, err := client.Do(req)
	if err != nil {
		return RecordDetail{}, errors.New("failed to request DNS Record: \n\t" + err.Error())
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return RecordDetail{}, errors.New("failed to read response: \n\t" + err.Error())
	}

	var recordResponse RecordDetailResponse
	err = json.Unmarshal(content, &recordResponse)
	if err != nil {
		return RecordDetail{}, errors.New("failed to decode response: \n\t" + err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return RecordDetail{}, errors.New("invalid status code (" + strconv.Itoa(res.StatusCode) + "): \n" + recordResponse.Error[0].ToDisplayString())
	}

	return recordResponse.Result, nil
}

func SetRecord(auth CFAuth, zoneID string, recordID string, record RecordDetail) error {
	client := &http.Client{}
	url := apiURL + "/" + zoneID + "/" + dnsRecordsPath + "/" + recordID

	payload, err := json.Marshal(record)
	if err != nil {
		return errors.New("failed to create request payload: \n\t" + err.Error())
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return errors.New("failed to build request: \n\t" + err.Error())
	}
	req.Header.Set(emailHeader, auth.Email)
	req.Header.Set(apiKeyHeader, auth.APIKey)
	req.Header.Set(contentTypeHeader, "application/json")

	res, err := client.Do(req)
	if err != nil {
		return errors.New("failed to request the update of DNS record: \n\t" + err.Error())
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("failed to read response: \n\t" + err.Error())
	}

	var recordResponse RecordDetailResponse
	err = json.Unmarshal(content, &recordResponse)
	if err != nil {
		return errors.New("failed to decode response: \n\t" + err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("invalid status code (" + strconv.Itoa(res.StatusCode) + "): \n" + recordResponse.Error[0].ToDisplayString())
	}

	return nil
}
