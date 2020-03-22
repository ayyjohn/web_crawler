package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("please specify a page to start on")
		os.Exit(1)
	} else {
		for _, arg := range args {
			fmt.Println(arg)
		}
	}
}
