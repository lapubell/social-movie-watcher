package main

import "strings"

var replacements map[string]string = map[string]string{
	":smile:": "😄",
	":D":      "😁",
	":)":      "😀",
	":hmm:":   "🤔",
}

func checkForEmojiTransform(s string) string {
	for check, replace := range replacements {
		s = strings.ReplaceAll(s, check, replace)
	}
	return s
}
