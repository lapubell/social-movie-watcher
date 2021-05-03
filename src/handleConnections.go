package main

import (
	"log"
	"net/http"
	"strings"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true
	participants[ws] = ""

	// now that we have registered the new connection, let's send the initial
	// state of the video
	// v.Video = "/video/video.mp4"
	ws.WriteJSON(map[string]interface{}{
		"initialMessage": true,
		"video":          v,
	})

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			delete(participants, ws)
			break
		}

		if msg.Message == "(just joined)" {
			participants[ws] = msg.Username
		}

		if strings.HasPrefix(msg.Message, "!!changeName:") {
			newName := parseNewNameFromMessage(msg.Message)
			participants[ws] = newName
			continue
		}

		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}
