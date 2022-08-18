package home

import (
	"strings"
)

func GetPingMessages(name string) string {
	s := []string{"This is the message you'll get,", name}

	v := strings.Join(s, " ")

	return v
}
