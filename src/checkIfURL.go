package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

// isValidUrl tests a string to determine if it is a well-structured url or not.
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func checkForImageTransform(s string) string {
	if !isValidUrl(s) {
		return s
	}

	res, err := http.Get(s)
	if err != nil {
		return s
	}

	if isImage(res.Header.Get("Content-Type")) {
		return "<img src='" + s + "' style='width: 100%' alt='Shared Image' />"
	}
	log.Printf("Type: %v", res.Header.Get("Content-Type"))

	return "<a href='" + s + "' target='_blank' rel='noopener noreferrer'>" + s + "</a>"
}

func isImage(mime string) bool {
	imageTypes := []string{
		"image/svg+xml",
		"image/gif",
		"image/jpeg",
		"image/png",
		"image/webp",
	}
	for _, t := range imageTypes {
		if strings.ToLower(mime) == strings.ToLower(t) {
			return true
		}
	}

	return false
}
