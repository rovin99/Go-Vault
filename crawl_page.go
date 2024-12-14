package main

import (
    "fmt"
    "net/url"
    "sync"
)

// config holds the configuration for the web crawler
type config struct {
    pages              map[string]int
    baseURL            *url.URL
    mu                 *sync.Mutex
    concurrencyControl chan struct{}
    wg                 *sync.WaitGroup
    maxPages           int
}

// addPageVisit adds a page visit to the pages map and returns whether it's the first visit
func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
    cfg.mu.Lock()
    defer cfg.mu.Unlock()
    if _, found := cfg.pages[normalizedURL]; found {
        cfg.pages[normalizedURL]++
        return false
    }
    cfg.pages[normalizedURL] = 1
    return true
}

// crawlPage crawls a page and its links

func (cfg *config) crawlPage(rawCurrentURL string) {
    defer func() {
        cfg.wg.Done()
        <-cfg.concurrencyControl
    }()

    currentURL, err := url.Parse(rawCurrentURL)
    if err != nil {
        fmt.Printf("invalid current URL: %v\n", err)
        return
    }

    if cfg.baseURL.Host != currentURL.Host {
        fmt.Printf("Skipping URL from different domain: %s\n", rawCurrentURL)
        return
    }

    normalizedURL := NormalizeURL(rawCurrentURL)
    if !cfg.addPageVisit(normalizedURL) {
        fmt.Printf("Already visited: %s\n", normalizedURL)
        return
    }

    html, err := getHTML(rawCurrentURL)
    if err != nil {
        fmt.Printf("error fetching HTML for %s: %v\n", rawCurrentURL, err)
        return
    }

    fmt.Printf("Crawling: %s\n", rawCurrentURL)

    urls, err := getURLsFromHTML(html, rawCurrentURL)
    if err != nil {
        fmt.Printf("error extracting URLs from %s: %v\n", rawCurrentURL, err)
        return
    }

    for _, u := range urls {
        normalizedU := NormalizeURL(u)
        cfg.mu.Lock()
        if _, found := cfg.pages[normalizedU]; found || len(cfg.pages) >= cfg.maxPages {
            cfg.mu.Unlock()
            continue
        }
        cfg.pages[normalizedU] = 0 // Mark as visited
        cfg.mu.Unlock()

        cfg.wg.Add(1)
        cfg.concurrencyControl <- struct{}{}
        go cfg.crawlPage(u)
    }
}