package restfulpath

import (
	. "bytes"
	"regexp"
)

// /users
// /users/:name
type Path struct {
	body string
}

func NewPath(body string) *Path {
	return &Path{body}
}

func (p *Path) Match(path string) (matched bool, params *map[string]string) {
	params = make(map[string]string)
	pa := Split(([]byte)(p.body), []byte{'/'})
	lenpa := len(pa)
	ta := Split(([]byte)(path), []byte{'/'})
	lenta := len(ta)
	if lenpa != lenta {
		matched = false
		return
	}
	for i, ip := range pa {
		ips := (string)(ip)
		its := (string)(ta[i])
		m, _ := regexp.Match("^:", ip)
		if m {
			params[ips[1:len(ips)]] = its
			continue
		}
		if ips != its {
			matched = false
			return
		}
	}
	matched = true

	return
}
