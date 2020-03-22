package main

import (
	"github.com/PuerkitoBio/goquery"
)

// ScrapeLinks parses all hrefs from a url and return them in a slice
func ScrapeLinks(uri string) []string {
	var links []string
	doc, _ := goquery.NewDocument(uri)
	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		links = append(links, href)
	})
	return links
}
