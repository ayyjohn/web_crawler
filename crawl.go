package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("please specify a page to start on")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("please only specify one page to start on")
	}

	retrieve(args[0])
}

func retrieve(uri string) {
	resp, err := http.Get(uri)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
