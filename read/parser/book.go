package parser

import (
	"regexp"
	"remfath.com/crawler/engine"
)

var bookListRe = regexp.MustCompile(`<div class="wrap">
<a href="([^"]+)" class="title" hidefocus="hidefocus" target="_blank">([^<]+)</a>
<div class="u-stargrade" itemprop="aggregateRating" itemscope itemtype="http://schema.org/AggregateRating">
<div class="icon grade10"></div>
<span [^<]+</span>
</div>
<div class="u-author">

<span>([^<]+)</span>

</div>
<p class="desc">([^<]+)</p>`)

func ParseBookList(contents []byte, category string) engine.ParseResult {
	match := bookListRe.FindAllStringSubmatch(string(contents), -1)

	var result engine.ParseResult
	for _, m := range match {
		book := engine.Book{Title: m[2], Author: m[3], Desc: m[4], Url: "http://www.duokan.com" + m[1], Category: category}
		result.Items = append(result.Items, book)
		result.Requests = append(result.Requests, engine.Request{Url: "http://www.duokan.com" + m[1], ParserFunc: engine.NilParser})
	}

	return result
}
