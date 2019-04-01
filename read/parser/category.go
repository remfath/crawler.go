package parser

import (
	"regexp"
	"remfath.com/crawler/engine"
)

var categoryListRe = regexp.MustCompile(`<li class="level1">
<div class="wrap">
<a href="([^"]+)" class="" hidefocus="hidefocus" >
<span>([^<]+)</span>
<em class="num">[\d]+</em>
<b></b>
</a>
</div>
</li>`)

func ParseCategoryList(contents []byte) engine.ParseResult {
	match := categoryListRe.FindAllStringSubmatch(string(contents), -1)

	var result engine.ParseResult
	for _, m := range match {
		category := m[2]
		result.Items = append(result.Items, category)
		req := engine.Request{Url: "http://www.duokan.com" + m[1], ParserFunc: func(c []byte) engine.ParseResult {
			return ParseBookList(c, category)
		}}
		result.Requests = append(result.Requests, req)
	}

	return result
}
