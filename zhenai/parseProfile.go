package zhenai

import (
	"github.com/markbest/crawler/engine"
	"regexp"
	"strconv"
	"strings"
)

var (
	idRe            = regexp.MustCompile(`<p class="brief-info fs14 lh32 c9f">ID：([^<]+)<span`)
	ageRe           = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	heightRe        = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	inComeRe        = regexp.MustCompile(`<td><span class="label">月收入：</span>(.+)元</td>`)
	marriageRe      = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	educationRe     = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	workCityRe      = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
	professionRe    = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	constellationRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
	nativePlaceRe   = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := &Profile{Url: url, Name: name}

	// profile id
	match := idRe.FindSubmatch(contents)
	if len(match) >= 2 {
		id, _ := strconv.Atoi(strings.TrimSpace(string(match[1])))
		profile.Id = id
	}

	// profile age
	match = ageRe.FindSubmatch(contents)
	if len(match) >= 2 {
		age, _ := strconv.Atoi(string(match[1]))
		profile.Age = age
	}

	// profile height
	match = heightRe.FindSubmatch(contents)
	if len(match) >= 2 {
		height, _ := strconv.Atoi(string(match[1]))
		profile.Height = height
	}

	// profile income
	match = inComeRe.FindSubmatch(contents)
	if len(match) >= 2 {
		profile.Income = string(match[1])
	}

	// profile marriage
	match = marriageRe.FindSubmatch(contents)
	if len(match) >= 2 {
		profile.Marriage = string(match[1])
	}

	// profile education
	match = educationRe.FindSubmatch(contents)
	if len(match) >= 2 {
		profile.Education = string(match[1])
	}

	// profile work city
	match = workCityRe.FindSubmatch(contents)
	if len(match) >= 2 {
		profile.WorkCity = string(match[1])
	}

	// profile profession
	match = professionRe.FindSubmatch(contents)
	if len(match) >= 2 {
		profile.Profession = string(match[1])
	}

	// profile constellation
	match = constellationRe.FindSubmatch(contents)
	if len(match) >= 2 {
		profile.Constellation = string(match[1])
	}

	// profile native place
	match = nativePlaceRe.FindSubmatch(contents)
	if len(match) >= 2 {
		profile.NativePlace = string(match[1])
	}
	result := engine.ParseResult{Items: []interface{}{profile}, Requests: nil}
	return result
}
