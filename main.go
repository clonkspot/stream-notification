package main

import (
	"log"
	"net/http"
	"os"
	"regexp"

	redis "github.com/fzzy/radix/extra/pool"
)

var twitchClientId = os.Getenv("TWITCH_CLIENT_ID")
var channels = os.Getenv("CHANNELS")
var port = os.Getenv("PORT")
var gameFilter = regexp.MustCompile(os.Getenv("GAME_FILTER"))
var redisNetwork = defaultValue(os.Getenv("REDIS_NETWORK"), "tcp")
var redisAddress = defaultValue(os.Getenv("REDIS_ADDRESS"), "127.0.0.1:6379")
var redisPool *redis.Pool

func defaultValue(val, def string) string {
	if val == "" {
		return def
	} else {
		return val
	}
}

func initRedis() {
	var err error
	redisPool, err = redis.NewPool(redisNetwork, redisAddress, 1)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initRedis()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			streamHandler(w, r)
		default:
			w.WriteHeader(400)
		}
	})

	port := os.Getenv("PORT")
	log.Print("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
