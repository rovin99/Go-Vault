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

	// Remove the scheme
	u.Scheme = ""

	// Remove the trailing slash if present
	normalizedPath := strings.TrimSuffix(u.Host + u.Path, "/")

	return normalizedPath
}
