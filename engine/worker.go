package engine

import (
	"time"
)

func Work(r Request) (ParseResult, error) {
	time.Sleep(time.Millisecond * 10)
	content, err := Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.ParseFunc(content), nil
}
