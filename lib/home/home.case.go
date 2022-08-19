package home

import (
	"strings"
)

func GetPingMessages(name string, locales string) string {
	var s string
	if locales == "en" {
		s = "This is the message you'll get,"
	} else {
		s = "Inilah pesan yang kamu dapatkan,"
	}
	msg := []string{s, name}
	v := strings.Join(msg, " ")

	return v
}
