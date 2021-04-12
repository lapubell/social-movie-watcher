package main

// Video the status of the video we are all watching
type Video struct {
	IsPlaying        bool   `json:"isPlaying"`
	CurrentTimestamp int    `json:"timestamp"`
	Video            string `json:"video"`
}
