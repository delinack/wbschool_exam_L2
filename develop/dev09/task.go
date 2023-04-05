package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
	"sync"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type result struct {
	url   string
	links []string
}

func main() {
	url := os.Args[1]
	results := make(chan result)
	var wg sync.WaitGroup

	wg.Add(1)
	go crawl(url, &wg, results)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("Fetched %s, found %d links\n", r.url, len(r.links))
	}
}

func crawl(url string, wg *sync.WaitGroup, results chan<- result) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Error parsing %s: %v\n", url, err)
		return
	}

	links := extractLinks(doc)
	results <- result{url: url, links: links}

	for _, link := range links {
		if strings.HasPrefix(link, "http") {
			wg.Add(1)
			go crawl(link, wg, results)
		}
	}
}

func extractLinks(n *html.Node) []string {
	var links []string

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, extractLinks(c)...)
	}

	return links
}
