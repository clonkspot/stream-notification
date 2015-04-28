package main

import (
	"bytes"
	"sync"
	"text/template"
)

var leagueUpdates chan []ResponseItem
var initLeagueOnce sync.Once
var leagueMotdTemplate struct {
	de *template.Template
	en *template.Template
}

func initLeague() {
	leagueUpdates = make(chan []ResponseItem)
	leagueMotdTemplate.de = template.Must(template.New("league_motd_de").Parse("Gerade live: Clonk auf Twitch\nMOTDURL={{.URL}}"))
	leagueMotdTemplate.en = template.Must(template.New("league_motd_en").Parse("Live right now: Clonk on Twitch\nMOTDURL={{.URL}}"))
	go updateLeague()
}

// Receives Twitch games from the `leagueUpdates` channel and puts them as MOTD
// in Redis.
// TODO: Cleanup on SIGTERM
func updateLeague() {
	r, err := redisPool.Get()
	if err != nil {
		panic(err)
	}

	rAppendStrings := func(op, key string, items []string) {
		args := make([]interface{}, len(items)+1)
		args[0] = key
		for k, v := range items {
			args[k+1] = v
		}
		r.Append(op, args...)
	}

	var prevDE, prevEN []string
	var streams []ResponseItem
	for {
		streams = <-leagueUpdates
		rAppendStrings("SREM", "league:motd:de", prevDE)
		rAppendStrings("SREM", "league:motd:en", prevEN)
		prevDE = prevDE[:0]
		prevEN = prevEN[:0]
		for _, stream := range streams {
			if stream.URL == "" {
				continue
			}
			var buf bytes.Buffer
			leagueMotdTemplate.de.Execute(&buf, stream)
			prevDE = append(prevDE, buf.String())
			buf.Reset()
			leagueMotdTemplate.en.Execute(&buf, stream)
			prevEN = append(prevEN, buf.String())
		}
		rAppendStrings("SADD", "league:motd:de", prevDE)
		rAppendStrings("SADD", "league:motd:en", prevEN)
		// Clear the pipeline.
		r.GetReply()
		r.GetReply()
		r.GetReply()
		r.GetReply()
	}
}

func UpdateLeague(streams []ResponseItem) {
	initLeagueOnce.Do(initLeague)
	leagueUpdates <- streams
}
