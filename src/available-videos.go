package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func availableVideos() []string {
	var files []string

	root := os.Getenv("VIDEO_FOLDER")
	log.Print("Video folder: " + root)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == root {
			return nil
		}
		output := strings.Replace(path, root+"/", "", 1)
		files = append(files, output[:len(output)-4])
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}
