package main

import (
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
