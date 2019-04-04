package parser

import (
	"regexp"
	"remfath.com/crawler/engine"
	"remfath.com/crawler/model"
	"strconv"
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
		categoryName := m[2]
		categoryId := getCategoryId(m[1])
		result.Items = append(result.Items, model.Category{Id: categoryId, Name: categoryName})
		req := engine.Request{Url: "http://www.duokan.com" + m[1], ParserFunc: func(c []byte) engine.ParseResult {
			return ParsePageList(c, categoryId, categoryName)
		}}
		result.Requests = append(result.Requests, req)
	}

	return result
}

var cidRe = regexp.MustCompile(`/list/([\d]+)-[\d]+`)

func getCategoryId(path string) int {
	match := cidRe.FindStringSubmatch(path)
	idStr := match[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	return id
}
