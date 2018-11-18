package payeer_api

import (
	"fmt"
	"testing"
)

func TestPayout(t *testing.T) {
	client := New("P000", "000", "000")

	res, err := client.GetPaySystems()
	if err != nil {
		t.Error(err.Error())
	} else {
		for _, item := range res.List {
			switch percent := item.CommissionSitePercent.(type) {
			case float64:
				fmt.Println("Это float64:",percent)
			case string:
				fmt.Println("Это string:",percent)
			}
		}
	}
}
