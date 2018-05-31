package zhenai

import (
	"github.com/markbest/crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]+>([^<]+)</a>`)

func ParseCityList(content []byte) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(content, -1)
	requests := make([]engine.Request, 0)
	for _, m := range matches {
		requests = append(requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return engine.ParseResult{Items: nil, Requests: requests}
}
