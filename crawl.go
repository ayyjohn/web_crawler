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
	} else if len(args) == 1 {
		for _, arg := range args {
			fmt.Println(arg)
		}
	} else {
		fmt.Println("please only specify one page to start on")
	}
}
