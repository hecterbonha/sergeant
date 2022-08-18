package home

import (
	"testing"
)

func TestGetPingMessages(t *testing.T) {
	name := "Gladys"
	msg := GetPingMessages(name)

	if msg != "This is the message you'll get, Gladys" {
		t.Error("Test Failed")
	}
}
