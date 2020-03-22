package main

import (
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeLinks parses all hrefs from a url and return them in a slice
func ScrapeLinks(uri string) []string {
	var links []string
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		return make([]string, 0)
	}
	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		links = append(links, href)
	})
	return links
}

// FixURL removes the query string from a URL and makes it absolute relative to the baseURL
func FixURL(href, base string) string {
	hrefURL, baseURL := parseOrEmpty(href), parseOrEmpty(base)
	urlWithoutQueryString := removeQueryString(hrefURL)
	absoluteURL := makeURLAbsolute(urlWithoutQueryString, baseURL)
	return absoluteURL.String()
}

func parseOrEmpty(uri string) *url.URL {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	return u
}

func makeURLAbsolute(href, base *url.URL) *url.URL {
	return base.ResolveReference(href)
}

func removeQueryString(url *url.URL) *url.URL {
	url.RawQuery = ""
	return url
}

// ConcurrentSet is a set suitable for concurrent situations that locks during reads and writes
type ConcurrentSet struct {
	set  map[string]bool
	lock sync.Mutex
}

// Add adds a value to the ConcurrentSet
func (s *ConcurrentSet) Add(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set[str] = true
}

// Contains returns true if the string is in the ConcurrentSet
func (s *ConcurrentSet) Contains(str string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.set[str]
}

// Length returns the number of items in the ConcurrentSet
func (s *ConcurrentSet) Length() int {
	return len(s.set)
}
