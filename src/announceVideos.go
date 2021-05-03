package main

import "time"

func announceVideos() {
	for {
		systemMessage := map[string]interface{}{
			"videos":    availableVideos(),
			"newStatus": "list",
		}
		sendToAll(systemMessage)

		time.Sleep(60 * time.Second)
	}
}
