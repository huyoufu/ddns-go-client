package dns

import (
	"strings"
)

type DomainStruct struct {
	RR   string
	Main string
}

// Analysis 解析域名成两部分 例如 www.jk1123.com ===> www jk1123.com
func Analysis(domain string) DomainStruct {
	if domain == "" {
		return DomainStruct{"", ""}
	}
	count := strings.Count(domain, ".")
	if count == 0 {
		return DomainStruct{"", ""}
	}
	if count == 1 {
		return DomainStruct{"", domain}
	}
	var r = []rune(domain)
	var c = 0
	var index = 0
	for i := len(r) - 1; i > 0; i-- {
		if string(r[i]) == "." {
			c++
			if c == 2 {
				index = i
				break
			}
		}
	}
	return DomainStruct{string(r[0:index]), string(r[index+1:])}
}
