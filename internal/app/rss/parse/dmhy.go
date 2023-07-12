package parse

import (
	"context"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/anguloc/zet/internal/app/rss/data"
	"github.com/anguloc/zet/pkg/log"
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

func (d *Dmhy) Run(ctx context.Context) (*data.List, error) {
	raw := &dmhyData{}
	if err := xml.Unmarshal(d.data, raw); err != nil {
		log.Error(ctx, "DmhyXmlParseErr", log.NamedError("err", err))
		return nil, err
	}
	res := &data.List{
		StartTime: time.Now(),
	}
	res.Data = make([]*data.Item, 0, len(raw.Channel.Item))

	for _, v := range raw.Channel.Item {
		res.Data = append(res.Data, &data.Item{
			Title:       v.Title,
			Link:        v.Link,
			Author:      v.Author,
			PubDate:     d.parsePubDate(v.PubDate),
			Description: v.Description,
			Bittorrent:  v.Enclosure.URL,
			Guid:        v.Guid.Text,
			UK:          d.parseUk(v.Guid.Text),
		})
	}
	return res, nil
}

func (d *Dmhy) parsePubDate(date string) time.Time {
	t, err := time.ParseInLocation(time.RFC1123Z, date, time.Local)
	if err != nil {
		return time.Time{}
	}
	if t.Location() != time.Local {
		t = t.Local()
	}
	return t
}

func (d *Dmhy) parseUk(str string) string {
	// 找字符串中的第一个数字
	res := ""
	for _, v := range str {
		s := string(v)
		_, ok := strconv.Atoi(s)
		if ok == nil {
			res += s
			continue
		}
		if res != "" {
			break
		}
	}
	return res
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
