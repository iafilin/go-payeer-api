package payeer_api

import (
	"bytes"
	"encoding/json"
)

type GetPaySystemsRes struct {
	Error
	List map[string]struct {
		ID                    string      `json:"id"`
		Name                  string      `json:"name"`
		GateCommission        interface{} `json:"gate_commission"`
		GateCommissionMin     interface{} `json:"gate_commission_min"`
		GateCommissionMax     interface{} `json:"gate_commission_max"`
		Currencies            []string    `json:"currencies"`
		CommissionSitePercent interface{} `json:"commission_site_percent"`
		RFields               map[string]struct {
			Name    string `json:"name"`
			RegExpr string `json:"reg_expr"`
			Example string `json:"example"`
		} `json:"r_fields"`
		SumMin interface{} `json:"sum_min"`
		SumMax interface{} `json:"sum_max"`
	} `json:"list"`
}

func (p *Payeer) GetPaySystems() (*GetPaySystemsRes, error) {
	data := &bytes.Buffer{}
	p.data.Add("action", "getPaySystems")
	_, err := data.WriteString(p.data.Encode())
	if err != nil {
		return nil, err
	}
	res, err := p.request(data)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	output := &GetPaySystemsRes{}
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
