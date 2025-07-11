package ip

import "testing"

func TestQuery(t *testing.T) {
	originalAPIURL := IpifyAPIURL

	IpifyAPIURL = "https://completlyfalseapi"

	_, err := Query()
	if err == nil {
		t.Error("Request to incorrect URL (" + IpifyAPIURL + ") succeded.")
	}

	IpifyAPIURL = originalAPIURL
}
