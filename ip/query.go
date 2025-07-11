// Package ip implements function to query public IP of the machine.
package ip

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

var IpifyAPIURL = "https://api4.ipify.org"

func Query() (string, error) {
	res, err := http.Get(IpifyAPIURL)
	if err != nil {
		return "", errors.New("Invalid status code received from " + IpifyAPIURL + ": 500")
	}

	if res.StatusCode != 200 {
		return "", errors.New("Invalid status code (" + strconv.Itoa(res.StatusCode) + ") from " + IpifyAPIURL)
	}
	ip, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("Invalid response from " + IpifyAPIURL + ": " + err.Error())
	}

	return string(ip), nil
}
