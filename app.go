package main

import (
	//"fmt"
	"github.com/dandyvica/gogrep/gg"
)

func main() {
	// fetch all options and flags
	options := gg.CliArgs()
	//fmt.Printf("%v\n", options)

	for _, element := range options.Files {
		gg.SearchIntoAFile(&element, &options)
	}
}
