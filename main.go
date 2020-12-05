package main

import (
	"flag"
	"github.com/prongbang/echogen/pkg/genx"
)

func main() {
	feature := flag.String("f", "", "-f=feature-name")
	flag.Parse()

	gen := genx.New()
	gen.Process(*feature)
}