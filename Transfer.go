package payeer_api

import (
	"bytes"
	"encoding/json"
	"errors"
)

type TransferResponse struct {
	Error
	HistoryID int    `json:"historyId"`
}

func (p *Payeer) Transfer(curIn, sum, curOut, to, comment string) (*TransferResponse, error) {
	data := &bytes.Buffer{}
	p.data.Add("action", "balance")
	p.data.Add("curIn", curIn)
	p.data.Add("sum", sum)
	p.data.Add("curOut", curOut)
	p.data.Add("to", to)
	p.data.Add("comment", comment)

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
