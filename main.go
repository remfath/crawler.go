package main

import (
	"remfath.com/crawler/engine"
	"remfath.com/crawler/read/parser"
)

func main() {
	url := "http://www.duokan.com/list/1-1"
	request := engine.Request{Url: url, ParserFunc: parser.ParseCategoryList}
	engine.Run(request)
}
