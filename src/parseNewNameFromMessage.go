package main

import "strings"

func parseNewNameFromMessage(blob string) string {
	chunks := strings.Split(blob, "NEW::")
	if len(chunks) != 2 {
		return "(Lurker)"
	}

	newNameChunk := strings.TrimSpace(chunks[1])
	if len(newNameChunk) < 2 {
		return "(Lurker)"
	}

	return strings.TrimSpace(newNameChunk[:len(newNameChunk)-2])
}
