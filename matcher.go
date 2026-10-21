package main

import "strings"

func shouldReply(content string) bool {
	return strings.Contains(strings.ToLower(content), "secret")
}
