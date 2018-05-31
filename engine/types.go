package engine

type ParserFunc func(contents []byte) ParseResult

type Request struct {
	Url       string
	ParseFunc ParserFunc
}

type ParseResult struct {
	Items    []interface{}
	Requests []Request
}