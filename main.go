package main

import (
	"flag"
	"github.com/prongbang/gestgen/utils"
)

func main() {
	feature := flag.String("f", "", "-f=feature-name")
	flag.Parse()

	fileGenerator := utils.NewFileGenerator()
	fileGenerator.GenerateAll(*feature)
}