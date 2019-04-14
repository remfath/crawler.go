package main

import (
	"github.com/remfath/crawler.go/engine"
	"github.com/remfath/crawler.go/persist"
	"github.com/remfath/crawler.go/read/parser"
	"github.com/remfath/crawler.go/scheduler"
)

func main() {
	url := "http://www.duokan.com/list/1-1"
	request := engine.Request{Url: url, ParserFunc: parser.ParseCategoryList}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
		ItemChan:  persist.ItemSaver(),
	}
	e.Run(request)
}
