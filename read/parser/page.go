package parser

import (
	"fmt"
	"regexp"
	"github.com/remfath/crawler.go/engine"
	"strconv"
)

var pageListRe = regexp.MustCompile(`<span>共([\d]+)页</span>`)

func ParsePageList(contents []byte, categoryId int, categoryName string) engine.ParseResult {
	match := pageListRe.FindStringSubmatch(string(contents))
	totalPage, err := strconv.Atoi(match[1])
	var result engine.ParseResult
	if err == nil {
		for page := 1; page <= totalPage; page++ {
			result.Items = append(result.Items, fmt.Sprintf("%d-%d", categoryId, page))
			req := engine.Request{Url: fmt.Sprintf("http://www.duokan.com/list/%d-%d", categoryId, page), ParserFunc: func(c []byte) engine.ParseResult {
				return ParseBookList(c, categoryName)
			}}
			result.Requests = append(result.Requests, req)
		}
	}

	return result
}
