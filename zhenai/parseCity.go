package zhenai

import (
	"github.com/markbest/crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*">([^<]*)</a>`)

func ParseCity(content []byte) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(content, -1)
	requests := make([]engine.Request, 0)
	for _, m := range matches {
		url := string(m[1])
		name := string(m[2])
		requests = append(requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, url, name)
			},
		})
	}
	return engine.ParseResult{Items: nil, Requests: requests}
}
