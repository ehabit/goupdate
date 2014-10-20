package main

import (
	"testing"
)

// func TestIsDir(t *testing.T) {

// }

// func TestParseFilenames(t *testing.T) {

// }

func TestIsGoFile(t *testing.T) {
	filename := "test.go"
	if !IsGoFile(filename) {
		t.Error("IsGoFile returned false for filename test.go")
	}
}

// func TestCheckDirForGo(t *testing.T) {

// }

// func TestUpdatePackage(t *testing.T) {

// }

// func TestUpdateCount(t *testing.T) {

// }
