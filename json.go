package main

// https://mholt.github.io/json-to-go/
type TwitchStreamsResponse struct {
	Total   int `json:"_total"`
	Streams []struct {
		Game        string  `json:"game"`
		Viewers     int     `json:"viewers"`
		AverageFps  float64 `json:"average_fps"`
		VideoHeight int     `json:"video_height"`
		CreatedAt   string  `json:"created_at"`
		ID          int64   `json:"_id"`
		Channel     struct {
			Mature                       bool   `json:"mature"`
			Status                       string `json:"status"`
			BroadcasterLanguage          string `json:"broadcaster_language"`
			DisplayName                  string `json:"display_name"`
			Game                         string `json:"game"`
			Delay                        int    `json:"delay"`
			Language                     string `json:"language"`
			ID                           int    `json:"_id"`
			Name                         string `json:"name"`
			CreatedAt                    string `json:"created_at"`
			UpdatedAt                    string `json:"updated_at"`
			Logo                         string `json:"logo"`
			Banner                       string `json:"banner"`
			VideoBanner                  string `json:"video_banner"`
			Background                   string `json:"background"`
			ProfileBanner                string `json:"profile_banner"`
			ProfileBannerBackgroundColor string `json:"profile_banner_background_color"`
			Partner                      bool   `json:"partner"`
			URL                          string `json:"url"`
			Views                        int    `json:"views"`
			Followers                    int    `json:"followers"`
			Links                        struct {
				Self          string `json:"self"`
				Follows       string `json:"follows"`
				Commercial    string `json:"commercial"`
				StreamKey     string `json:"stream_key"`
				Chat          string `json:"chat"`
				Features      string `json:"features"`
				Subscriptions string `json:"subscriptions"`
				Editors       string `json:"editors"`
				Teams         string `json:"teams"`
				Videos        string `json:"videos"`
			} `json:"_links"`
		} `json:"channel"`
		Preview struct {
			Small    string `json:"small"`
			Medium   string `json:"medium"`
			Large    string `json:"large"`
			Template string `json:"template"`
		} `json:"preview"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"streams"`
	Links struct {
		Summary  string `json:"summary"`
		Followed string `json:"followed"`
		Next     string `json:"next"`
		Featured string `json:"featured"`
		Self     string `json:"self"`
	} `json:"_links"`
}

type ResponseItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
