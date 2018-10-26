package parser

import (
	"regexp"

	"../../engine"
)

const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "user:"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProile(c, name)
			},
		})
	}
	return result
}