package payeer

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
	AuthError string   `json:"auth_error"`
	Errors    []string `json:"errors"`
}

func (a *Error) Error() string {
	return strings.Join(a.Errors,"\n")
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
	resData := &Error{}
	if err := json.NewDecoder(res.Body).Decode(resData); err != nil {
		return err
	}
	if len(resData.Errors) != 0 {
		return errors.New(resData.Errors[0])
	} else {
		return nil
	}
}
