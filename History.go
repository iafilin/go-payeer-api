package payeer_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

type HistoryType string

const (
	HistoryTypeIncoming = HistoryType("incoming")
	HistoryTypeOutgoing = HistoryType("outgoing")
)

type HistorySort string

const (
	HistorySortAsc  = HistorySort("asc")
	HistorySortDesc = HistorySort("desc")
)

type HistoryItem struct {
	ID               string `json:"id"`
	Date             string `json:"date"`
	Type             string `json:"type"`
	Status           string `json:"status"`
	From             string `json:"from"`
	To               string `json:"to"`
	CreditedAmount   string `json:"creditedAmount"`
	CreditedCurrency string `json:"creditedCurrency"`
	Protect          string `json:"protect"`
	ProtectDay       int    `json:"protectDay"`
	Comment          string `json:"comment"`
}
type HistoryResponse struct {
	Error
	Params struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Sort  string `json:"sort"`
		Count int    `json:"count"`
	} `json:"params"`
	History map[string]HistoryItem `json:"history,null"`
}

func (p *Payeer) History(count int, historyType HistoryType, historySort HistorySort) (*HistoryResponse, error) {
	data := &bytes.Buffer{}
	p.data.Add("action", "history")
	p.data.Add("count", string(count))
	p.data.Add("type", string(historyType))
	p.data.Add("sort", string(historySort))

	_, err := data.WriteString(p.data.Encode())
	if err != nil {
		return nil, err
	}
	res, err := p.request(data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resData := &HistoryResponse{}
	if err := json.NewDecoder(res.Body).Decode(resData); err != nil {
		if strings.Contains(err.Error(), "HistoryResponse.history") {
			resData.History = map[string]HistoryItem{}
		} else {
			return nil, err
		}

	}
	if len(resData.Error.Error()) != 0 {
		return nil, errors.New(resData.Error.Error())
	} else {
		return resData, err
	}
}
