package payeer_api

import (
	"bytes"
	"encoding/json"
)

type GetBalanceRes struct {
	Error
	Balance struct {
		BCH struct {
			BUDGET       string `json:"BUDGET"`
			DOSTUPNO     string `json:"DOSTUPNO"`
			DOSTUPNOSYST string `json:"DOSTUPNO_SYST"`
		} `json:"BCH"`
		BTC struct {
			BUDGET       string `json:"BUDGET"`
			DOSTUPNO     string `json:"DOSTUPNO"`
			DOSTUPNOSYST string `json:"DOSTUPNO_SYST"`
		} `json:"BTC"`
		DAA struct {
			BUDGET       string `json:"BUDGET"`
			DOSTUPNO     string `json:"DOSTUPNO"`
			DOSTUPNOSYST string `json:"DOSTUPNO_SYST"`
		} `json:"DAA"`
		ETH struct {
			BUDGET       string `json:"BUDGET"`
			DOSTUPNO     string `json:"DOSTUPNO"`
			DOSTUPNOSYST string `json:"DOSTUPNO_SYST"`
		} `json:"ETH"`
		EUR struct {
			BUDGET       string `json:"BUDGET"`
			DOSTUPNO     string `json:"DOSTUPNO"`
			DOSTUPNOSYST string `json:"DOSTUPNO_SYST"`
		} `json:"EUR"`
		LTC struct {
			BUDGET       string `json:"BUDGET"`
			DOSTUPNO     string `json:"DOSTUPNO"`
			DOSTUPNOSYST string `json:"DOSTUPNO_SYST"`
		} `json:"LTC"`
		RUB struct {
			BUDGET       string `json:"BUDGET"`
			DOSTUPNO     string `json:"DOSTUPNO"`
			DOSTUPNOSYST string `json:"DOSTUPNO_SYST"`
		} `json:"RUB"`
		USD struct {
			BUDGET       string `json:"BUDGET"`
			DOSTUPNO     string `json:"DOSTUPNO"`
			DOSTUPNOSYST string `json:"DOSTUPNO_SYST"`
		} `json:"USD"`
	} `json:"balance"`
}

func (p *Payeer) GetBalance() (*GetBalanceRes, error) {
	data := &bytes.Buffer{}
	p.data.Add("action", "balance")
	_, err := data.WriteString(p.data.Encode())
	if err != nil {
		return nil, err
	}
	res, err := p.request(data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	output := &GetBalanceRes{}
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
