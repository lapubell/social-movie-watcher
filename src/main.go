package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
var v Video                                  // video state machine

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message object
type Message struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

// Video the status of the video we are all watching
type Video struct {
	IsPlaying        bool `json:"isPlaying"`
	CurrentTimestamp int  `json:"timestamp"`
}

func main() {
	// instantiate a video state machine
	v = Video{
		CurrentTimestamp: 0,
		IsPlaying:        false,
	}

	// update the video timestamp in a go routine
	go incrementVideoTimestamp()

	// massive video files need to be broken up into smaller bits for easier streaming
	http.HandleFunc("/video/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../video/video.mp4")
	})

	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8081 and log any errors
	log.Println("http server started on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

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

	// now that we have registered the new connection, let's send the initial
	// state of the video
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
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
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

func incrementVideoTimestamp() {
	for {
		if v.IsPlaying {
			v.CurrentTimestamp++
		}
		time.Sleep(time.Second * 1)
	}
}
