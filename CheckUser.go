package payeer_api

import (
	"bytes"
	"encoding/json"
	"errors"
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
		return nil, err
	}
	if len(resData.Error.Error()) != 0 {
		return nil, errors.New(resData.Error.Error())
	} else {
		return resData, err
	}
}
