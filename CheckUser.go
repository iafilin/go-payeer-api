package payeer_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

type CheckUserResponse struct {
	Error
	HistoryID int `json:"historyId"`
}

func (p *Payeer) CheckUser(accountNumber string) (*TransferResponse, error) {
	data := &bytes.Buffer{}
	p.data.Add("action", "checkUser")
	p.data.Add("user", accountNumber)
	_, err := data.WriteString(p.data.Encode())
	if err != nil {
		return nil, err
	}
	res, err := p.request(data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resData := &TransferResponse{}
	if err := json.NewDecoder(res.Body).Decode(resData); err != nil {
		if strings.Contains(err.Error(),"errors") {
			resData.Errors = []string{}
		}else{
			return nil, err
		}
	}
	if len(resData.Errors) != 0 {
		return nil, errors.New(resData.Error.Errors[0])
	} else {
		return resData, err
	}
}
