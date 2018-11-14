package payeer

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

type RFieldsItem struct {
	Name    string `json:"name"`
	RegExpr string `json:"reg_expr"`
	Example string `json:"example"`
}
type ListItem struct {
	ID                    string                 `json:"id"`
	Name                  string                 `json:"name"`
	GateCommission        map[string]string      `json:"gate_commission"`
	GateCommissionMin     map[string]string      `json:"gate_commission_min"`
	GateCommissionMax     map[string]string      `json:"gate_commission_max"`
	Currencies            []string               `json:"currencies"`
	CommissionSitePercent string                 `json:"commission_site_percent"`
	RFields               map[string]RFieldsItem `json:"r_fields"`
	SumMin                map[string]string      `json:"sum_min"`
	SumMax                map[string]string      `json:"sum_max"`
}
type GetPaySystemsResponse struct {
	Error
	List map[string]ListItem `json:"list"`
}

func (p *Payeer) GetPaySystems() (*GetPaySystemsResponse, error) {
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
	structData := &GetPaySystemsResponse{}
	var mapData map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&mapData); err != nil {
		return nil, err
	}
	decode(mapData, structData)
	if len(structData.Errors) != 0 {
		return nil, errors.New(structData.Error.Errors[0])
	} else {
		return structData, err
	}
}
func decode(mapData map[string]interface{}, response *GetPaySystemsResponse) {
	response.AuthError = mapData["auth_error"].(string)
	response.List = map[string]ListItem{}
	for _, err := range mapData["errors"].([]interface{}) {
		response.Errors = append(response.Errors, err.(string))
	}

	for k, v := range mapData["list"].(map[string]interface{}) {
		item := v.(map[string]interface{})
		response.List[k] = ListItem{
			ID:   item["id"].(string),
			Name: item["name"].(string),
			GateCommission: func() map[string]string {
				if value, ok := item["gate_commission"].(map[string]string); ok {
					return value
				} else {
					return nil
				}
			}(),
			GateCommissionMin: func() map[string]string {
				if value, ok := item["gate_commission_min"].(map[string]string); ok {
					return value
				} else {
					return nil
				}
			}(),
			GateCommissionMax: func() map[string]string {
				if value, ok := item["gate_commission_max"].(map[string]string); ok {
					return value
				} else {
					return nil
				}
			}(),
			Currencies: func() []string {
				var values []string
				for _, v := range item["currencies"].([]interface{}) {
					values = append(values, v.(string))
				}
				return values
			}(),
			CommissionSitePercent: func() string {
				if value, ok := item["commission_site_percent"].(string); ok {
					return value
				} else {
					return strconv.FormatFloat(item["commission_site_percent"].(float64), 'f', 2, 64)
				}
			}(),
			RFields: func() map[string]RFieldsItem {
				var retData = map[string]RFieldsItem{}
				for k, item := range item["r_fields"].(map[string]interface{}) {
					a := item.(map[string]interface{})
					retData[k] =  RFieldsItem{
						Name: func() string {
							if val, ok := a["name"].(string); ok {
								return val
							} else {
								return ""
							}
						}(),
						RegExpr: func() string {
							if val, ok := a["reg_expr"].(string); ok {
								return val
							} else {
								return ""
							}
						}(),
						Example: func() string {
							if val, ok := a["example"].(string); ok {
								return val
							} else {
								return ""
							}
						}(),
					}
				}
				return retData
			}(),
			SumMin: func() map[string]string {
				if value, ok := item["sum_min"].(map[string]string); ok {
					return value
				} else {
					return nil
				}
			}(),
			SumMax: func() map[string]string {
				if value, ok := item["sum_max"].(map[string]string); ok {
					return value
				} else {
					return nil
				}
			}(),
		}
	}
}
