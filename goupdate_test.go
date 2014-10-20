package main

import (
	"os"
	"testing"
)

func TestIsDir(t *testing.T) {
	// setUp
	os.Mkdir("tmp", 0777)

	currentDirectory, err := os.Getwd()
	if err != nil {
		t.Error("An error occured trying to get current working directory:", err)
	}

	// test
	if !IsDir(currentDirectory + "/tmp") {
		t.Error("IsDir returned false for directory tmp")
	}

	// tearDown
	err2 := os.Remove("tmp")
	if err2 != nil {
		t.Error("An error occured trying to remove tmp directory:", err2)
	}
}

func TestIsGoFile(t *testing.T) {
	filename := "test.go"
	if !IsGoFile(filename) {
		t.Error("IsGoFile returned false for filename test.go")
	}
}

// func TestParseFilenames(t *testing.T) {

// }

// func TestCheckDirForGo(t *testing.T) {

// }

// func TestUpdatePackage(t *testing.T) {

// }

// func TestUpdateCount(t *testing.T) {

// }
