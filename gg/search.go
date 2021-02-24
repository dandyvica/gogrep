package gg

import (
	"bufio"
	"fmt"
	"os"
)

func SearchIntoAFile(fileName *string, options *CliOptions) {
	// this will hold current context
	var ctx Context
	ctx.fileName = fileName

	// open file fo reading and test errors
	file, err := os.Open(*fileName)

	if err != nil {
		if !options.ignoreErrors {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		return
	}

	// uses a new scanner
	fscanner := bufio.NewScanner(file)

	// use this buffer to get lines
	for fscanner.Scan() {
		ctx.currentLine = fscanner.Text()
		ctx.lineNumber++

		// test line against the regexp
		ctx.Grep(options)
	}

	// number of matches at the end of the search
	if options.matchesNumber {
		ctx.DisplayMatches()
	}

	// wen only wnat file names
	if options.onlyFiles {
		ctx.DisplayFileNames()
	}

}
