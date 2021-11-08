package main

import "strings"

func checkForGiphy(s string) string {
	if len(s) < 7 {
		return s
	}

	lowerS := strings.ToLower(s)

	if lowerS[0:7] != ":giphy " {
		return s
	}

	tag := s[7:]
	gifImg := getRandomGifUrlByTag(tag)
	return gifImg
}
