package engine

import (
	"log"
	"remfath.com/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s...\n", r.Url)

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetch %s error: %s\n", r.Url, err)
			continue
		}

		result := r.ParserFunc(body)
		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			log.Printf("Got items: %#v\n", item)
		}
	}
}
