//go:build integration
// +build integration

package battery_test

import (
	"battery"
	"testing"
)

func TestGetPmsetOutput(t *testing.T) {
	t.Parallel()
	text, err := battery.GetPmsetOutput()
	if err != nil {
		t.Fatal(err)
	}
	status, err := battery.ParsePmsetOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Charge: %d%%", status.ChargePercent)
}
