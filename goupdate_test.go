package main

import (
	"os"
	"reflect"
	"testing"
)

func TestIsDir(t *testing.T) {
	// setUp
	os.Mkdir("tmp", 0777)

	// test
	if !IsDir("tmp") {
		t.Error("IsDir returned false for directory tmp")
	}

	// cleanUp
	err := os.Remove("tmp")
	if err != nil {
		t.Error("An error occured trying to remove tmp/,", err)
	}
}

func TestIsGoFile(t *testing.T) {
	filename := "test.go"
	if !IsGoFile(filename) {
		t.Error("IsGoFile returned false for filename test.go")
	}
}

func TestParseFilenames(t *testing.T) {
	// setUp
	os.Mkdir("tmp", 0777)

	os.Create("tmp/test.txt")
	// os.Create("tmp/test2.md")

	test_data := []string{"test.txt"}

	currentDirectory, err := os.Getwd()
	if err != nil {
		t.Error("An error occured trying to get current working directory,", err)
	}

	filenames := ParseFilenames(currentDirectory + "/tmp")
	if !reflect.DeepEqual(filenames, test_data) {
		t.Error("ParseFilenames did not return expected value")
	}

	// cleanUp
	err = os.Remove("tmp/test.txt")
	if err != nil {
		t.Error("An error occured trying to remove tmp/test.txt,", err)
	}

	err = os.Remove("tmp")
	if err != nil {
		t.Error("An error occured trying to remove tmp/,", err)
	}
}

// func TestCheckDirForGo(t *testing.T) {

// }

// func TestUpdatePackage(t *testing.T) {

// }

// func TestGetUpdateCount(t *testing.T) {

// }
