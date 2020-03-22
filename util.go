package main

import (
	"net/url"

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

// FixURL takes a link and makes it absolute if it isn't
func FixURL(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseURL, err := url.Parse(base)
	if err != nil {
		return ""
	}

	uri = baseURL.ResolveReference(uri)
	return uri.String()
}
