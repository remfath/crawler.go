package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/remfath/crawler.go/engine"
)

var Items map[string]struct{}

func init() {
	Items = make(map[string]struct{})
}

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			respId, err := save(item)
			if err != nil {
				fmt.Printf("ERR: %s\n", err)
			} else {
				fmt.Printf("SUC: %s\n", respId)
			}
		}
	}()
	return out
}

func save(item engine.Item) (string, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index(item.Index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
