package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Ensure the current URL is on the same domain as the base URL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("invalid base URL: %v\n", err)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("invalid current URL: %v\n", err)
		return
	}

	if baseURL.Host != currentURL.Host {
		return
	}

	// Normalize the current URL
	normalizedURL := NormalizeURL(rawCurrentURL)

	// Check if the page has already been crawled
	if count, found := pages[normalizedURL]; found {
		pages[normalizedURL] = count + 1
		return
	}

	// Mark the page as visited
	pages[normalizedURL] = 1

	// Fetch the HTML content
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching HTML for %s: %v\n", rawCurrentURL, err)
		return
	}

	fmt.Printf("Crawling: %s\n", rawCurrentURL)

	// Extract URLs from the HTML
	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		fmt.Printf("error extracting URLs from %s: %v\n", rawCurrentURL, err)
		return
	}

	// Recursively crawl each URL found on the page
	for _, u := range urls {
		crawlPage(rawBaseURL, u, pages)
	}
}
