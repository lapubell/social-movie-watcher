package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading in ENV file, using defaults")
	}

	// instantiante the default video struct
	v = Video{
		Video: "/video/video.mp4",
	}

	screeningRooms = make([]room, 1)

	// update the video timestamp in a go routine
	go incrementVideoTimestamp()

	// massive video files need to be broken up into smaller bits for easier streaming
	http.HandleFunc("/video/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache, max-age=0")

		filename := strings.Replace(r.URL.String(), "/video/", "", 1)
		filename, err := url.QueryUnescape(filename)
		_, err = os.Stat(os.Getenv("VIDEO_FOLDER") + filename)
		if os.IsNotExist(err) {
			fmt.Println(os.Getenv("VIDEO_FOLDER") + filename)
			http.ServeFile(w, r, "../video/video.mp4")
		}
		http.ServeFile(w, r, os.Getenv("VIDEO_FOLDER")+filename)
	})

	// Create a simple file server
	fs := noCacheStaticAssets(http.FileServer(http.Dir("../public")))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	numRutinesString := os.Getenv("SIZE")
	numRutines, err := strconv.Atoi(numRutinesString)
	if err != nil {
		numRutines = 1
	}
	for i := 0; i < numRutines; i++ {
		go handleMessages()
	}

	// Start the server on port 8081 (by default) and log any errors
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Println("http server started on :" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
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
