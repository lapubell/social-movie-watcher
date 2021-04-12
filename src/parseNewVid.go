package main

import "strings"

func parseNewVid(messageMessage string) string {
	prefixStrip := strings.Replace(messageMessage, "!!change:", "", 1)

	if strings.HasSuffix(prefixStrip, "!!") {
		return prefixStrip[:len(prefixStrip)-2]
	}

	return prefixStrip
}
