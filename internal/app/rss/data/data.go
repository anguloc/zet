package data

import "time"

type Item struct {
	Title       string
	Link        string
	Author      string
	PubDate     time.Time
	Description string
	Bittorrent  string
	Guid        string
	UK          string
}

type List struct {
	StartTime time.Time
	Data      []*Item
}

type TransmissionItem struct{}

type TransmissionList struct {
	StartTime time.Time
	Data      []*TransmissionItem
}
