package payeer_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const apiEndpoint = "https://payeer.com/ajax/api/api.php"

type Payeer struct {
	accountNumber string
	apiId         string
	apiKey        string
	data          *url.Values
}

type Error struct {
	AuthError string      `json:"auth_error"`
	Errors    interface{} `json:"errors"`
}

func (p *Error) Error() string {
	switch err := p.Errors.(type) {
	case bool:
		return ""
	case []string:
		return strings.Join(err, ", ")
	case string:
		return err
	default:
		return ""
	}

}

func New(accountNumber, apiId, apiKey string) *Payeer {
	data := &url.Values{}
	data.Add("account", accountNumber)
	data.Add("apiId", apiId)
	data.Add("apiPass", apiKey)
	return &Payeer{
		accountNumber: accountNumber,
		apiId:         apiId,
		apiKey:        apiKey,
		data:          data,
	}
}
func (p *Payeer) getUrl(apiMethod string) string {
	return fmt.Sprintf("%s?%s", apiEndpoint, apiMethod)
}

func (p *Payeer) request(data *bytes.Buffer) (*http.Response, error) {
	req, err := http.NewRequest("POST", apiEndpoint, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		return res, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Payeer.com server error: %d", res.StatusCode))
	}
}

func (p *Payeer) CheckAuth() error {
	data := &bytes.Buffer{}
	_, err := data.WriteString(p.data.Encode())
	if err != nil {
		return err
	}
	res, err := p.request(data)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var resData = make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&resData); err != nil {
		return err
	}
	switch resData["errors"].(type) {
	case []string:
		return errors.New(resData["errors"].([]string)[0])
	default:
		return nil
	}
}
