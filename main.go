package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
)

var twitchClientId = os.Getenv("TWITCH_CLIENT_ID")
var channels = os.Getenv("CHANNELS")
var port = os.Getenv("PORT")
var gameFilter = regexp.MustCompile(os.Getenv("GAME_FILTER"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			streamHandler(w, r)
		default:
			w.WriteHeader(400)
		}
	})

	port := os.Getenv("PORT")
	log.Print("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
