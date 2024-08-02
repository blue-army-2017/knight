package main

import "time"

type LegacyData struct {
	Member map[string]MemberData `json:"member"`
	Season map[string]SeasonData `json:"season"`
}

type MemberData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Active    bool   `json:"active"`
}

type SeasonData struct {
	Name    string                    `json:"name"`
	Created time.Time                 `json:"created"`
	Games   map[string]SeasonGameData `json:"games"`
}

type SeasonGameData struct {
	Opponent       string    `json:"opponent"`
	Date           time.Time `json:"date"`
	Home           bool      `json:"home"`
	Mode           string    `json:"mode"`
	PresentMembers []string  `json:"presentMembers"`
}
