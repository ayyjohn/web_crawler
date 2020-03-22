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
