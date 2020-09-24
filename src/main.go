package main

import (
	"flag"
	"fmt"
)

func main() {
	var builder string
	flag.StringVar(&builder, "builder", "maven", "maven or gradle")

	flag.Parse()

	fmt.Println(builder)
}