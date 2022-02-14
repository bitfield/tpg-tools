package main

import (
	"battery"
	"fmt"
	"os"
)

func main() {
	status, err := battery.GetStatus()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't read battery status: %v", err)
	}
	fmt.Printf("Battery %d%% charged\n", status.ChargePercent)
}
