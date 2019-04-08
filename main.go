package main

import (
	"remfath.com/crawler.go/engine"
	"remfath.com/crawler.go/read/parser"
	"remfath.com/crawler.go/scheduler"
)

func main() {
	url := "http://www.duokan.com/list/1-1"
	request := engine.Request{Url: url, ParserFunc: parser.ParseCategoryList}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
	}
	e.Run(request)
}
