package parse

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/anguloc/zet/internal/pkg/log"
)

type Dmhy struct {
	data []byte
}

func NewDmhy() *Dmhy {
	return &Dmhy{}
}

func (d *Dmhy) SetData(data []byte) error {
	d.data = data
	return nil
}

func (d *Dmhy) Run(ctx context.Context) (*Result, error) {
	data := &dmhyData{}
	if err := xml.Unmarshal(d.data, data); err != nil {
		fmt.Println(err)
		return nil, err
		log.Error(ctx, "xmlParseErr", log.NamedError("err", err))
	}
	println(len(data.Channel.Item))
	for _, v := range data.Channel.Item {
		println(v.Title)
	}
	res, _ := json.MarshalIndent(data.Channel.Item[0], "", "\t")
	println(string(res))
	fmt.Printf("%+v", data.Channel.Item[0])
	return nil, fmt.Errorf("a")
}

type dmhyData struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Content string   `xml:"content,attr"`
	Wfw     string   `xml:"wfw,attr"`
	Channel struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		PubDate     string `xml:"pubDate"`
		Item        []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Description string `xml:"description"`
			Enclosure   struct {
				Text   string `xml:",chardata"`
				URL    string `xml:"url,attr"`
				Length string `xml:"length,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"enclosure"`
			Author string `xml:"author"`
			Guid   struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Category struct {
				Text   string `xml:",chardata"`
				Domain string `xml:"domain,attr"`
			} `xml:"category"`
		} `xml:"item"`
	} `xml:"channel"`
}
