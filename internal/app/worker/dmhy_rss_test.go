package worker

import (
	"context"
	"fmt"
	"testing"

	"github.com/anguloc/zet/internal/app/repo/kv"
)

func TestA(t *testing.T) {
	ctx := context.Background()
	w := &DmhyRss{}

	url := "https://www.dmhy.org/topics/rss/rss.xml"

	body, code, err := w.queryRss(ctx, url)
	if err != nil {
		fmt.Println(err)
		return
	}

	cache := kv.New()
	cache.Put("dmhy", body)

	println(code)

}
