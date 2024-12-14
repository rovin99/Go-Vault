# Web Crawler

A simple web crawler written in Go that fetches URLs from a given base URL and visualizes the crawling results.

## Getting Started

### Prerequisites

* Go (version 1.16 or later)

### Cloning the Repository

To clone this repository, run the following commands:

```bash
git clone https://github.com/yourusername/web-crawler.git
cd web-crawler
```

### Running the Crawler

To run the crawler, use the following command:

```bash
go run main.go <URL> <maxConcurrency> <maxPages>
```

* `<URL>`: The base URL to start crawling from
* `<maxConcurrency>`: The maximum number of concurrent requests
* `<maxPages>`: The maximum number of pages to crawl

#### Example

```bash
go run main.go https://example.com 5 100
```

### Generating Graph Visualization

The crawler will generate a graph visualization of the crawled pages and save it as `graph.png`.

## License

This project is licensed under the MIT License - see the LICENSE file for details.


