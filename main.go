package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/distatus/battery"
)

func main() {
	logFile := flag.String("logFile", "./battery.log", "Log file path. Default: ./battery.log")
	flag.Parse()

	fh := getLogFile(*logFile)
	defer fh.Close()

	getBatteryStatus()
	os.Exit(0)
}

func getLogFile(logFile string) *os.File {
	fmt.Printf("Log file path: %s\n", logFile)
	fh, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)

	if err != nil {
		fmt.Printf("Error opening log file: %s\n", err)
		os.Exit(1)
	}

	return fh
}

func getBatteryStatus() {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		os.Exit(2)
	}
	for i, battery := range batteries {
		fmt.Printf("Bat%d:\n", i)
		fmt.Printf("state: %s,\n", battery.State)
		fmt.Printf("current capacity: %.2f mWh,\n", battery.Current)
		fmt.Printf("last full capacity: %.2f mWh,\n", battery.Full)
		fmt.Printf("design capacity: %.2f mWh,\n", battery.Design)
		fmt.Printf("charge rate: %.2f mW,\n", battery.ChargeRate)
		fmt.Printf("voltage: %.2f V,\n", battery.Voltage)
		fmt.Printf("design voltage: %.2f V\n", battery.DesignVoltage)
	}

}
