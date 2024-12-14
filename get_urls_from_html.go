package main

import (
	
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
    if htmlBody == "" {
        return []string{}, nil
    }

    urls := []string{}
    urlSet := make(map[string]bool)

    baseURL, err := url.Parse(rawBaseURL)
    if err!= nil {
        return nil, err
    }

    htmlReader := strings.NewReader(htmlBody)
    doc, err := html.Parse(htmlReader)
    if err!= nil {
        return nil, err
    }

    var traverse func(*html.Node)
    traverse = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    href := attr.Val
                    _, err := url.Parse(href)
                    if err!= nil {
                        continue
                    }
                    resolvedURL, err := url.Parse(href)
                    if err!= nil {
                        continue
                    }
                    if resolvedURL.IsAbs() {
                        urlStr := resolvedURL.String()
                        if!urlSet[urlStr] {
                            urls = append(urls, urlStr)
                            urlSet[urlStr] = true
                        }
                    } else {
                        resolvedURL = baseURL.ResolveReference(resolvedURL)
                        urlStr := resolvedURL.String()
                        if!urlSet[urlStr] {
                            urls = append(urls, urlStr)
                            urlSet[urlStr] = true
                        }
                    }
                }
            }
        }
        for c := n.FirstChild; c!= nil; c = c.NextSibling {
            traverse(c)
        }
    }

    traverse(doc)
    return urls, nil
}