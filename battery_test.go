package main

import (
	"os"
	"path"
	"testing"
)

// TestOpenLogFile tests if log file is created
func TestOpenLogFile(t *testing.T) {
	testFilePath := path.Join(os.TempDir(), "gotest.log")
	_ = getLogFile(testFilePath)
	_, err := os.Stat(testFilePath)
	if os.IsNotExist(err) {
		t.Errorf("Error")
	}

	os.Remove(testFilePath)
}
