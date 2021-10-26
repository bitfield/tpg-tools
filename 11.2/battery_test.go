package battery_test

import (
	"battery"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToJSON(t *testing.T) {
	t.Parallel()
	batt := battery.Battery{
		Name:             "InternalBattery-0",
		ID:               10813539,
		ChargePercent:    100,
		TimeToFullCharge: "0:00",
		Present:          true,
	}
	wantBytes, err := os.ReadFile("testdata/battery.json")
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantBytes)
	got := batt.ToJSON()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
