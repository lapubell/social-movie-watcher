package main

import (
	"sort"
	"time"
)

func announceChatRoomPeople() []string {
	output := []string{}
	for _, name := range participants {
		output = append(output, name)
	}
	sort.Strings(output)
	return output
}

func broadcastChatRoomPeople() {
	for {
		systemMessage := map[string]interface{}{
			"peeps":     announceChatRoomPeople(),
			"newStatus": "peeps",
		}
		sendToAll(systemMessage)

		time.Sleep(time.Second * 5)
	}
}
