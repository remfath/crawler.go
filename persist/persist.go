package persist

import (
	"fmt"
	"github.com/remfath/crawler.go/model"
)

var Items map[string]struct{}

func init() {
	Items = make(map[string]struct{})
}

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			if book, ok := item.(model.Book); ok {
				if _, ok := Items[book.Title]; !ok {
					fmt.Printf("%s\n", book.Title)
					Items[book.Title] = struct{}{}
				}
			}
		}
	}()
	return out
}
