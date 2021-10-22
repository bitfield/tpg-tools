package battery_test

import (
	"battery"
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJSONOutput(t *testing.T) {
	t.Parallel()
	input := battery.Status{
		DrawingFrom: "AC Power",
		Batteries: []battery.Battery{
			{
				Name:             "InternalBattery-0",
				ID:               10813539,
				ChargePercent:    100,
				TimeToFullCharge: "0:00",
				Present:          true,
			},
		},
	}
	wantBytes, err := os.ReadFile("testdata/pmset.json")
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantBytes)
	buf := &bytes.Buffer{}
	err = input.WriteJSONTo(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
