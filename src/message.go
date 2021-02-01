package main

import "log"

// Message object
type Message struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		if msg.Message == "!!video-play!!" {
			log.Printf("video now playing: %v", v)
			v.IsPlaying = true
			systemMessage := map[string]interface{}{
				"video":     v,
				"newStatus": "play",
			}
			sendToAll(systemMessage)
			continue
		}

		if msg.Message == "!!video-pause!!" {
			v.IsPlaying = false
			v.CurrentTimestamp = msg.Timestamp
			systemMessage := map[string]interface{}{
				"video":     v,
				"newStatus": "pause",
			}
			sendToAll(systemMessage)
			continue
		}

		if msg.Message == "!!available-videos!!" {
			systemMessage := map[string]interface{}{
				"videos": availableVideos(),
			}
			sendToAll(systemMessage)
			continue
		}

		// default, send the chatted message to all clients
		sendToAll(msg)
	}
}

func sendToAll(thing interface{}) {
	// Send it out to every client that is currently connected
	for client := range clients {
		err := client.WriteJSON(thing)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
