package main

import (
	"log"
	"strings"
)

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

		if msg.Message == "!!who-dis!!" {
			systemMessage := map[string]interface{}{
				"peeps":     announceChatRoomPeople(),
				"newStatus": "peeps",
			}
			sendToAll(systemMessage)
			continue
		}

		if msg.Message == "!!available-videos!!" {
			systemMessage := map[string]interface{}{
				"videos":    availableVideos(),
				"newStatus": "list",
			}
			sendToAll(systemMessage)
			continue
		}

		if strings.HasPrefix(msg.Message, "!!change:") {
			v.CurrentTimestamp = 0
			v.IsPlaying = true
			newFile := "/video/" + parseNewVid(msg.Message)
			v.Video = newFile

			systemMessage := map[string]interface{}{
				"video":     v,
				"changeTo":  newFile,
				"newStatus": "change",
			}
			sendToAll(systemMessage)
			continue
		}

		msg.Message = checkForImageTransform(msg.Message)
		msg.Message = checkForEmojiTransform(msg.Message)
		msg.Message = checkForGiphy(msg.Message)

		// default, send the chatted message to all clients
		sendToAll(msg)
	}
}

func sendToAll(thing interface{}) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

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
