package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplate = "../../templates"
	os.Exit(m.Run())
}
