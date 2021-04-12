package main

import "testing"

func Test_weCanGetANewVideoFileFromAMessage(t *testing.T) {
	message := "!!change:Test"
	if parseNewVid(message) != "Test" {
		t.Error("Expected Test, got: ", parseNewVid(message))
	}
	message = "!!change:Test!!"
	if parseNewVid(message) != "Test" {
		t.Error("Expected Test, got: ", parseNewVid(message))
	}
}
