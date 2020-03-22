package main

import (
	"flag"
	"fmt"
	"os"
)

var visitedURLs = ConcurrentSet{set: make(map[string]bool)}

func main() {
	args := parseArgs()

	queue := make(chan string)

	go func() {
		queue <- args[0]
	}()

	for uri := range queue {
		enqueueLinks(uri, queue)
	}
}

func enqueueLinks(uri string, queue chan string) {
	if !visitedURLs.Contains(uri) {
		fmt.Println("fetching", uri)
		visitedURLs.Add(uri)
		links := ScrapeLinks(uri)
		for _, link := range links {
			absoluteLink := FixURL(link, uri)
			if uri != "" && !visitedURLs.Contains(absoluteLink) {
				go func(l string) { queue <- l }(absoluteLink)
			}
		}
	}
}

func retrieve(uri string) {
	fmt.Println(ScrapeLinks(uri))
}

func parseArgs() []string {
	flag.Parse()
	args := flag.Args()

	validateArgs(args)
	return args
}

func validateArgs(args []string) {
	if len(args) != 1 {
		fmt.Println("please specify exactly one page to start on")
		os.Exit(1)
	}
}
