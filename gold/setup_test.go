package main

import (
	"os"
	"testing"
)

var testApp app_config

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
