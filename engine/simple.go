package engine

import (
	"fmt"
	"log"
	"github.com/remfath/crawler.go/fetcher"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}

		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Got items: %#v\n", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s...\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, fmt.Errorf("Fetch %s error: %s\n", r.Url, err)
	}

	return r.ParserFunc(body), nil
}
