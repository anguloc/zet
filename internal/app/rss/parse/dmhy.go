package parse

import (
	"context"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/anguloc/zet/internal/app/rss"
	"github.com/anguloc/zet/internal/pkg/log"
)

type Dmhy struct {
	data []byte
}

func NewDmhy() *Dmhy {
	return &Dmhy{}
}

func (d *Dmhy) SetData(data []byte) IParse {
	d.data = data
	return d
}

func (d *Dmhy) Run(ctx context.Context) (*rss.List, error) {
	data := &dmhyData{}
	if err := xml.Unmarshal(d.data, data); err != nil {
		log.Error(ctx, "DmhyXmlParseErr", log.NamedError("err", err))
		return nil, err
	}
	res := &rss.List{
		StartTime: time.Now().Unix(),
	}
	res.Data = make([]*rss.Item, 0, len(data.Channel.Item))

	for _, v := range data.Channel.Item {
		println(v.Title)
		res.Data = append(res.Data, &rss.Item{
			Title: v.Title,
		})
	}
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
