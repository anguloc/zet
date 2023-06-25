package rss

type Item struct{}

type List struct {
	StartTime int64
	Data      []*Item
}

type TransmissionItem struct{}

type TransmissionList struct {
	StartTime int64
	Data      []*TransmissionItem
}
