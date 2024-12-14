package main

import (
    "fmt"
    "os"
    "strconv"
    "sync"
    "net/url"
    "sort"
)

type page struct {
    URL  string
    Count int
}


func sortPages(pages map[string]int) []page {
    var sortedPages []page
    for url, count := range pages {
        sortedPages = append(sortedPages, page{URL: url, Count: count})
    }

    // Sort the pages by count in descending order and then by URL in ascending order
    sort.Slice(sortedPages, func(i, j int) bool {
        if sortedPages[i].Count == sortedPages[j].Count {
            return sortedPages[i].URL < sortedPages[j].URL
        }
        return sortedPages[i].Count > sortedPages[j].Count
    })

    return sortedPages
}

func printReport(pages map[string]int, baseURL string) {
    fmt.Println("=============================")
    fmt.Printf("  REPORT for %s\n", baseURL)
    fmt.Println("=============================")

    sortedPages := sortPages(pages)
    for _, page := range sortedPages {
        fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
    }
}

// func main() {
//     args := os.Args[1:]
//     if len(args) != 3 {
//         fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
//         os.Exit(1)
//     }

//     baseURL, err := url.Parse(args[0])
//     if err != nil {
//         fmt.Printf("invalid base URL: %v\n", err)
//         os.Exit(1)
//     }

//     maxConcurrency, err := strconv.Atoi(args[1])
//     if err != nil || maxConcurrency <= 0 {
//         fmt.Println("invalid maxConcurrency value")
//         os.Exit(1)
//     }

//     maxPages, err := strconv.Atoi(args[2])
//     if err != nil || maxPages <= 0 {
//         fmt.Println("invalid maxPages value")
//         os.Exit(1)
//     }

//     cfg := &config{
//         pages:              make(map[string]int),
//         baseURL:            baseURL,
//         mu:                 &sync.Mutex{},
//         concurrencyControl: make(chan struct{}, maxConcurrency),
//         wg:                 &sync.WaitGroup{},
//         maxPages:           maxPages,
//     }

//     fmt.Printf("starting crawl of: %s\n", args[0])
//     cfg.wg.Add(1)
//     cfg.concurrencyControl <- struct{}{}
//     go cfg.crawlPage(args[0])
//     cfg.wg.Wait()

//     printReport(cfg.pages, args[0])
    
// } 

func main() {
    args := os.Args[1:]
    if len(args)!= 3 {
        fmt.Println("usage:./crawler URL maxConcurrency maxPages")
        os.Exit(1)
    }

    baseURL, err := url.Parse(args[0])
    if err!= nil {
        fmt.Printf("invalid base URL: %v\n", err)
        os.Exit(1)
    }

    maxConcurrency, err := strconv.Atoi(args[1])
    if err!= nil || maxConcurrency <= 0 {
        fmt.Println("invalid maxConcurrency value")
        os.Exit(1)
    }

    maxPages, err := strconv.Atoi(args[2])
    if err!= nil || maxPages <= 0 {
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

    printReport(cfg.pages, args[0])
}