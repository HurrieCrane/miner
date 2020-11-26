package utils

import (
	"net/url"
	"regexp"
)

var httpRegex = regexp.MustCompile("(https)")

func IsHTTPURL(uri string) bool {
	uuri, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	return httpRegex.Match([]byte(uuri.Scheme))
}
