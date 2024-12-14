package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"net/url"
	

)

func main() {
    args := os.Args[1:]
    if len(args) != 3 {
        fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
        os.Exit(1)
    }

    baseURL, err := url.Parse(args[0])
    if err != nil {
        fmt.Printf("invalid base URL: %v\n", err)
        os.Exit(1)
    }

    maxConcurrency, err := strconv.Atoi(args[1])
    if err != nil || maxConcurrency <= 0 {
        fmt.Println("invalid maxConcurrency value")
        os.Exit(1)
    }

    maxPages, err := strconv.Atoi(args[2])
    if err != nil || maxPages <= 0 {
        fmt.Println("invalid maxPages value")
        os.Exit(1)
    }

    cfg := &config{
        pages:              make(map[string]int),
        baseURL:            baseURL,
        mu:                 &sync.Mutex{},
        concurrencyControl: make(chan struct{}, maxConcurrency),
        wg:                 &sync.WaitGroup{},
        maxPages:           maxPages,
    }

    fmt.Printf("starting crawl of: %s\n", args[0])
    cfg.wg.Add(1)
    cfg.concurrencyControl <- struct{}{}
    go cfg.crawlPage(args[0])
    cfg.wg.Wait()

    fmt.Println("Crawled pages:")
    for page, count := range cfg.pages {
        fmt.Printf("%s: %d\n", page, count)
    }
} 