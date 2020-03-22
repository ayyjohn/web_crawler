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
	linkURL, linkErr := url.Parse(href)
	baseURL, baseErr := url.Parse(base)
	if linkErr != nil || baseErr != nil {
		return ""
	}

	urlWithoutQueryString := removeQueryString(linkURL)
	absoluteURL := makeURLAbsolute(urlWithoutQueryString, baseURL)
	return absoluteURL.String()
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
	lock sync.RWMutex
}

// Add adds a value to the ConcurrentSet
func (s *ConcurrentSet) Add(str string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set[str] = true
}

// Contains returns true if the string is in the ConcurrentSet
func (s *ConcurrentSet) Contains(str string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, ok := s.set[str]
	return ok
}

// Length returns the length of the underlying map in the ConcurrentSet
func (s *ConcurrentSet) Length() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.set)
}
