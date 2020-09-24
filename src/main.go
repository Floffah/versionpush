package main

import (
	"flag"
	"fmt"
)

func main() {
	var builder string
	flag.StringVar(&builder, "builder", "br", "maven or gradle")

	fmt.Println(builder)
}