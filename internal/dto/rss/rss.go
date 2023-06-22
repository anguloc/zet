package rss

type Item struct {
	Title string `json:"title"`
}

type List struct {
	Data []*Item
}
