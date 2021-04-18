package main

import "net/url"

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

	return "<a href='" + s + "' target='_blank' rel='noopener noreferrer'>" + s + "</a>"
}
