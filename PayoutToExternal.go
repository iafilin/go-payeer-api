package payeer_api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type OutputRes struct {
	Error
	OutputParams struct {
		SumIn  string `json:"sumIn"`
		CurIn  string `json:"curIn"`
		CurOut string `json:"curOut"`
		Ps     int    `json:"ps"`
		SumOut string `json:"sumOut"`
	} `json:"outputParams"`
}

func (p *Payeer) Output(ps, sumIn, curIn, curOut string, fields map[string]string) (*OutputRes, error) {
	data := &bytes.Buffer{}
	p.data.Add("action", "output")
	p.data.Add("ps", ps)
	p.data.Add("sumIn", sumIn)
	p.data.Add("curIn", curIn)
	p.data.Add("curOut", curOut)
	for paramKey, paramValue := range fields {
		p.data.Add(fmt.Sprintf("param_%s", paramKey), paramValue)
	}

	_, err := data.WriteString(p.data.Encode())
	if err != nil {
		return nil, err
	}
	res, err := p.request(data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	output := &OutputRes{}
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
