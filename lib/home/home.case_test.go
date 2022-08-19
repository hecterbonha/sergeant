package home

import (
	"testing"
)

func TestGetPingMessages(t *testing.T) {
	name := "Gladys"
	msgEN := GetPingMessages(name, "en")
	msgID := GetPingMessages(name, "id")

	if msgEN != "This is the message you'll get, Gladys" {
		t.Error("Test Failed")
	}
	if msgID != "Inilah pesan yang kamu dapatkan, Gladys" {
		t.Error("Test Failed")
	}
}
