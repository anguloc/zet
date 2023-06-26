package data

type Item struct {
	Title       string
	Link        string
	Author      string
	PubDate     string
	Description string
	Bittorrent  string
	Guid        Guid
}

type Guid struct {
	Text        string
	IsPermaLink bool
}

type List struct {
	StartTime int64
	Data      []*Item
}

type TransmissionItem struct{}

type TransmissionList struct {
	StartTime int64
	Data      []*TransmissionItem
}
