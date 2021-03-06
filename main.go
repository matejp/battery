package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/distatus/battery"
)

type batteryData struct {
	name           string
	value          string
	valueAsFloat64 float64
	unit           string
}

func main() {
	logFile := flag.String("logFile", "./battery.log", "Log file path. Default: ./battery.log")
	flag.Parse()

	fh := getLogFile(*logFile)
	defer fh.Close()

	data := getBatteryStatus()
	capacityPercentage := (data["Bat0"]["current capacity"].valueAsFloat64 / data["Bat0"]["full capacity"].valueAsFloat64) * 100
	// fmt.Printf("%.2f%%\n", capacityPercentage)

	fmt.Printf("(%s) state: %s %s => %.2f%% of full capacity: %s %s\n", "Bat0", data["Bat0"]["state"].value, data["Bat0"]["current capacity"].value, capacityPercentage, data["Bat0"]["full capacity"].value, data["Bat0"]["full capacity"].unit)

	os.Exit(0)
}

func getLogFile(logFile string) *os.File {
	// fmt.Printf("Log file path: %s\n", logFile)
	fh, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)

	if err != nil {
		fmt.Printf("Error opening log file: %s\n", err)
		os.Exit(1)
	}

	return fh
}

func getBatteryStatus() map[string]map[string]batteryData {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		os.Exit(2)
	}
	// fmt.Println(batteries)
	// initialized outer map
	batteryStats := make(map[string]map[string]batteryData)

	for i, battery := range batteries {
		// initialized inner map
		batteryNumber := fmt.Sprintf("Bat%d", i)
		batteryStats[batteryNumber] = make(map[string]batteryData)

		batteryStats[batteryNumber]["state"] =
			batteryData{name: "state", value: fmt.Sprintf("%s", battery.State), valueAsFloat64: 0.0, unit: ""}
		batteryStats[batteryNumber]["current capacity"] =
			batteryData{name: "current capacity", value: fmt.Sprintf("%.f2", battery.Current), valueAsFloat64: battery.Current, unit: "mWh"}
		batteryStats[batteryNumber]["full capacity"] =
			batteryData{name: "full capacity", value: fmt.Sprintf("%.f2", battery.Full), valueAsFloat64: battery.Full, unit: "mWh"}
		batteryStats[batteryNumber]["design capacity"] =
			batteryData{name: "design capacity", value: fmt.Sprintf("%.f2", battery.Design), valueAsFloat64: battery.Design, unit: "mWh"}
		batteryStats[batteryNumber]["charge rate"] =
			batteryData{name: "charge rate", value: fmt.Sprintf("%.f2", battery.ChargeRate), valueAsFloat64: battery.ChargeRate, unit: "mW"}
		batteryStats[batteryNumber]["voltage"] =
			batteryData{name: "voltage", value: fmt.Sprintf("%.f2", battery.Voltage), valueAsFloat64: battery.Voltage, unit: "V"}
		batteryStats[batteryNumber]["design voltage"] =
			batteryData{name: "design voltage", value: fmt.Sprintf("%.f2", battery.DesignVoltage), valueAsFloat64: battery.DesignVoltage, unit: "V"}

		// fmt.Printf("Bat%d:\n", i)
		// // fmt.Printf("state: %s,\n", battery.State)
		// fmt.Printf("current capacity: %.2f mWh,\n", battery.Current)
		// fmt.Printf("last full capacity: %.2f mWh,\n", battery.Full)
		// fmt.Printf("design capacity: %.2f mWh,\n", battery.Design)
		// fmt.Printf("charge rate: %.2f mW,\n", battery.ChargeRate)
		// fmt.Printf("voltage: %.2f V,\n", battery.Voltage)
		// fmt.Printf("design voltage: %.2f V\n", battery.DesignVoltage)
	}

	return batteryStats
}
