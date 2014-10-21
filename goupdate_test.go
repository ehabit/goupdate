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
		t.Error("IsDir() returned false for directory tmp")
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
		t.Error("IsGoFile() returned false for filename test.go")
	}
}

func TestParseFilenames(t *testing.T) {
	// setUp
	os.Mkdir("tmp", 0777)

	os.Create("tmp/test.txt")
	os.Create("tmp/test2.md")

	test_data := []string{"test.txt", "test2.md"}

	currentDirectory, err := os.Getwd()
	if err != nil {
		t.Error("An error occured trying to get current working directory,", err)
	}

	// test
	filenames := ParseFilenames(currentDirectory + "/tmp")
	if !reflect.DeepEqual(filenames, test_data) {
		t.Error("ParseFilenames() did not return expected value.")
	}

	// cleanUp
	err = os.Remove("tmp/test.txt")
	if err != nil {
		t.Error("An error occured trying to remove tmp/test.txt,", err)
	}

	err = os.Remove("tmp/test2.md")
	if err != nil {
		t.Error("An error occured trying to remove tmp/test2.md,", err)
	}

	err = os.Remove("tmp")
	if err != nil {
		t.Error("An error occured trying to remove tmp/,", err)
	}
}

func TestCheckDirForGo(t *testing.T) {
	// setUp
	os.Mkdir("tmp", 0777)

	os.Create("tmp/test.go")

	currentDirectory, err := os.Getwd()
	if err != nil {
		t.Error("An error occured trying to get current working directory,", err)
	}

	// test
	if !CheckDirForGo(currentDirectory + "/tmp") {
		t.Error("CheckDirForGo() did not properly identify a directory with a .go file in it.")
	}

	// cleanUp
	err = os.Remove("tmp/test.go")
	if err != nil {
		t.Error("An error occured trying to remove tmp/test.go,", err)
	}

	err = os.Remove("tmp")
	if err != nil {
		t.Error("An error occured trying to remove tmp/,", err)
	}
}

// Wercker seems to remove version control when it clones packages in its default build step
// This means wercker errors on go get -u with version control not found.
// The git clone goupdatebuild step in wercker.yml combined with the ssh-key for git access
// solve this problem and are required for TestUpdatePackage() to function properly on wercker
func TestUpdatePackage(t *testing.T) {

	GOPATH := os.Getenv("GOPATH")
	packPath := GOPATH + "/src/github.com/ehabit/goupdate"
	updated := UpdatePackage(packPath)
	if !updated {
		t.Error("An error occured trying to update this Go package,", packPath)
	}
}
