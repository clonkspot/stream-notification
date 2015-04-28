package main

import (
	"encoding/json"
	"net/http"
)

const CACHE_DURATION = 120

// Requests Twitch streams, answering with a filtered set.
func streamHandler(w http.ResponseWriter, r *http.Request) {
	// Always return JSON.
	w.Header().Set("Content-Type", "application/json")

	// Try getting a cached response.
	if cached := GetCache("streams"); cached != nil {
		w.Write(cached)
		return
	}

	sr, err := GetTwitchStreams(channels)
	if err != nil {
		w.Write([]byte("[]"))
		return
	}

	resitems := FilterStreams(sr)
	go UpdateLeague(resitems)

	b, err := json.Marshal(resitems)
	if err != nil || len(resitems) == 0 {
		b = []byte("[]")
	}
	w.Write(b)
	PutCache("streams", CACHE_DURATION, b)
}
