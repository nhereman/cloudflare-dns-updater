package cloudflare

import "strconv"

type ErrorResponse struct {
	Code             int    `json:"code"`
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
	Source           struct {
		Pointer string `json:"pointer"`
	} `json:"source"`
}

func (e ErrorResponse) ToDisplayString() string {
	res := ""
	res += "\tcode: " + strconv.Itoa(e.Code) + "\n"
	res += "\tmessage: " + e.Message + "\n"
	res += "\tdocumentation_url: " + e.DocumentationURL + "\n"
	res += "\tsource: \n"
	res += "\t\tpointer: " + e.Source.Pointer + "\n"
	return res
}
