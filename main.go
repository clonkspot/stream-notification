package main

import (
	"encoding/json"
	"github.com/franela/goreq"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

var twitchClientId = os.Getenv("TWITCH_CLIENT_ID")
var channels = os.Getenv("CHANNELS")
var port = os.Getenv("PORT")
var gameFilter = regexp.MustCompile(os.Getenv("GAME_FILTER"))

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// Always return JSON.
	w.Header().Set("Content-Type", "application/json")

	query := url.Values{}
	query.Add("channel", channels)
	req := goreq.Request{
		Uri:         "https://api.twitch.tv/kraken/streams",
		QueryString: query,
	}
	req.AddHeader("Client-ID", twitchClientId)
	res, err := req.Do()
	if err != nil {
		w.Write([]byte("[]"))
		return
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	var sr TwitchStreamsResponse
	if err = dec.Decode(&sr); err != nil {
		w.Write([]byte("[]"))
		return
	}

	var resitems []ResponseItem
	for _, stream := range sr.Streams {
		if !gameFilter.MatchString(stream.Game) {
			continue
		}
		resitems = append(resitems, ResponseItem{Name: stream.Channel.DisplayName, URL: stream.Channel.URL})
	}

	b, err := json.Marshal(resitems)
	if err != nil || len(resitems) == 0 {
		w.Write([]byte("[]"))
		return
	}
	w.Write(b)
}

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
