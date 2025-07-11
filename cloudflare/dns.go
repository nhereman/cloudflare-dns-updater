// Package cloudflare provides utilities to contact cloudflare API
package cloudflare

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type RecordDetailResponse struct {
	Error  []ErrorResponse `json:"errors"`
	Result struct {
		Domain string `json:"name"`
		IP     string `json:"content"`
	} `json:"result"`
}

type RecordDetail struct {
	Domain string
	IP     string
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

	return RecordDetail{
		Domain: recordResponse.Result.Domain,
		IP:     recordResponse.Result.IP,
	}, nil
}
