package payeer_api

import (
	"bytes"
	"encoding/json"
)

func (p *Payeer) CheckUser(accountNumber string) (error) {
	data := &bytes.Buffer{}
	p.data.Add("action", "checkUser")
	p.data.Add("user", accountNumber)
	_, err := data.WriteString(p.data.Encode())
	if err != nil {
		return err
	}
	res, err := p.request(data)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	output := &Error{}
	if err := json.NewDecoder(res.Body).Decode(output); err != nil {
		return err
	}

	switch e := output.Errors.(type) {
	case []string:
		if len(e) > 0 {
			return output
		} else {
			return nil
		}
	default:
		return nil
	}

}
