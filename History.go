package payeer_api

import (
	"bytes"
	"encoding/json"
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

type HistoryRes struct {
	Error
	Params struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Sort  string `json:"sort"`
		Count int    `json:"count"`
	} `json:"params"`
	History map[string]struct {
		ID               string `json:"id"`
		Date             string `json:"date"`
		Type             string `json:"type"`
		Status           string `json:"status"`
		To               string `json:"to"`
		CreditedAmount   string `json:"creditedAmount"`
		CreditedCurrency string `json:"creditedCurrency"`
		PaySystem        string `json:"paySystem"`
	} `json:"history"`
}

func (p *Payeer) History(count int, historyType HistoryType, historySort HistorySort) (*HistoryRes, error) {
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
	output := &HistoryRes{}
	if err := json.NewDecoder(res.Body).Decode(output); err != nil {
		return nil, err
	}

	switch e := output.Errors.(type) {
	case []string:
		if len(e) > 0 {
			return nil, &output.Error
		} else {
			return output, nil
		}
	default:
		return output, nil
	}
}
