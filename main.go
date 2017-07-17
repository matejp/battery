package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	logFile := flag.String("logFile", "./battery.log", "Log file path. Default: ./battery.log")
	flag.Parse()

	fmt.Printf("Log file path: %s\n", *logFile)
	fh, err := os.OpenFile(*logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	defer fh.Close()

	if err != nil {
		fmt.Printf("Error opening log file: %s\n", err)
		os.Exit(1)
	}


	os.Exit(0)
}
