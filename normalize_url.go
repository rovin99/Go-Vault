package main

import (
	"net/url"
	"strings"
)


func NormalizeURL(rawURL string) string {
    u, err := url.Parse(rawURL)
    if err != nil {
        return rawURL
    }
    u.Path = strings.TrimSuffix(u.Path, "/")
    return u.String()
}