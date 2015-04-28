package main

import (
	"encoding/json"
	"github.com/franela/goreq"
	"net/url"
)

// Requests streams from the given channels (comma-delimited).
func GetTwitchStreams(channels string) (sr TwitchStreamsResponse, err error) {
	query := url.Values{}
	query.Add("channel", channels)
	req := goreq.Request{
		Uri:         "https://api.twitch.tv/kraken/streams",
		QueryString: query,
	}
	req.AddHeader("Client-ID", twitchClientId)
	res, err := req.Do()
	if err != nil {
		return TwitchStreamsResponse{}, err
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&sr); err != nil {
		return TwitchStreamsResponse{}, err
	}
	return
}
