package battery

import (
	"encoding/json"
)

type Battery struct {
	Name             string
	ID               int64
	ChargePercent    int
	TimeToFullCharge string
	Present          bool
}

func (b Battery) ToJSON() string {
	output, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(output)
}
