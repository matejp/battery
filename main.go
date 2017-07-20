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
	// fmt.Print(data)
	// for i := range data {
	// 	for j := range data[i] {
	// 		fmt.Printf(" (%s) %s => %s\n", i, j, data[i][j])
	// 	}
	// }
	// capacityPercentage, err := strconv.ParseFloat(data["Bat0"]["current capacity float"], 32)
	// if err != nil {
	// 	fmt.Println("Error parsing data.", err)
	// 	os.Exit(2)
	// }
	// designCapacityPercentage, err := strconv.ParseFloat(data["Bat0"]["design capacity float"], 32)
	// if err != nil {
	// 	fmt.Println("Error parsing data.", err)
	// 	os.Exit(2)
	// }

	// fmt.Printf("%.2f%% (%s)", capacityPercentage/designCapacityPercentage*100, "test")
	fmt.Printf("(%s) %s/%s => %s/%s\n", "Bat0", "current capacity", "design capacity", data["Bat0"][0].value, data["Bat0"][1].value)

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

func getBatteryStatus() map[string][]batteryData {
	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		os.Exit(2)
	}
	batteryStats := make(map[string][]batteryData)

	for i, battery := range batteries {
		batteryStats[fmt.Sprintf("Bat%d", i)] =
			append(batteryStats[fmt.Sprintf("Bat%d", i)],
				batteryData{name: "state", value: fmt.Sprintf("%s", battery.State), valueAsFloat64: 0.0, unit: ""})
		batteryStats[fmt.Sprintf("Bat%d", i)] =
			append(batteryStats[fmt.Sprintf("Bat%d", i)],
				batteryData{name: "current capacity", value: fmt.Sprintf("%.f2", battery.Current), valueAsFloat64: battery.Current, unit: "mWh"})
		// "state":              fmt.Sprintf("%s", battery.State),
		// "state value unit":   "string",
		// "current capacity":   fmt.Sprintf("%.2f", battery.Current), //mWh
		// "last full capacity": fmt.Sprintf("%.2f", battery.Full),

		// "current capacity":       fmt.Sprintf("%.2f mWh (%.2f V)", battery.Current, battery.Voltage),
		// "current capacity float": fmt.Sprintf("%.2f", battery.Current),
		// "last full capacity":     fmt.Sprintf("%.2f mWh", battery.Full),
		// "design capacity":        fmt.Sprintf("%.2f mWh (%.2f V)", battery.Design, battery.DesignVoltage),
		// "design capacity float":  fmt.Sprintf("%.2f", battery.Design),
		// }

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
