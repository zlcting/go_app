package parser

import (
	"regexp"
	"strconv"

	"../../engine"
	"../../model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)

var Marriage = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

func ParseProile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, Marriage)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(content []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
