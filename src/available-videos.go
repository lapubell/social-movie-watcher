package main

import (
	"os"
	"path/filepath"
	"strings"
)

func availableVideos() []string {
	var files []string

	root := os.Getenv("VIDEO_FOLDER")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == root {
			return nil
		}
		output := strings.Replace(path, root, "", 1)
		if output != ".gitkeep" && output != "video.mp4" {
			files = append(files, output)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}
