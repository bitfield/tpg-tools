package battery

import (
	"encoding/json"
	"io"
)

type Status struct {
	DrawingFrom string
	Batteries   []Battery
}

type Battery struct {
	Name             string
	ID               int64
	ChargePercent    int
	TimeToFullCharge string
	Present          bool
}

func (s Status) WriteJSONTo(w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(s)
}
