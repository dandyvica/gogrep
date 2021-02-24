package gg

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

const Usage = `
NAME
	ggr: this is a grep utility written in Go. Project repository: https://github.com/dandyvica/gogrep
	Files matching the regex (see regexp Go syntax: https://golang.org/pkg/regexp/syntax) are displayed in
	addition to matching lines, with different colors.

USAGE
	ggr [OPTIONS...] PATTERN [FILE...]

OPTIONS
	-c, -count
		don't show matching lines but only matching count

	-e, -ignore-errors
		silently ignore I/O errors when opening files (e.g.: due to file permissions)

	-h, -no-filename
		only print matching lines, suppress file name output

	-i, -ignore-case
		ignore case when searching

	-l, -files-with-matches
		only print files matching the pattern

	-n, -line-number
		print line number in addition to other data

	-o, -only-matching
		print only tokens matching the regex

	-v, -invert-match
		invert the matching by adding the ignore case modifier (?i)

EXAMPLES
`

// this will hold all options
type CliOptions struct {
	// compiled regex from command line
	re *regexp.Regexp

	// list of files (Unix) or pattern (Windows)
	Files []string

	// we want the line number to be displayed
	requestLinesNumbers bool

	// we want the file name to be displayed
	noFileName bool

	// we want to ignore io errors
	ignoreErrors bool

	// case insensitive
	ignoreCase bool

	// invert search
	invertMatch bool

	// we want the number of matches
	matchesNumber bool

	// we only want file names
	onlyFiles bool

	// we only want maching tokens, not the whole line
	onlyMatching bool
}

func CliArgs() CliOptions {
	// init struct
	var options CliOptions

	// if set, we want the line number from the file
	flag.BoolVar(&options.requestLinesNumbers, "n", false, "line number is printed out")
	flag.BoolVar(&options.requestLinesNumbers, "line-number", false, "line number is printed out")

	// if set, we want the file name
	flag.BoolVar(&options.noFileName, "h", false, "file name is not printed out")
	flag.BoolVar(&options.noFileName, "no-filename", false, "file name is not printed out")

	// if set, we ignore error
	flag.BoolVar(&options.ignoreErrors, "e", false, "ignore error we opening files")
	flag.BoolVar(&options.ignoreErrors, "ignore-errors", false, "ignore error we opening files")

	// if set, no case sensitive
	flag.BoolVar(&options.ignoreCase, "i", false, "ignore case we searching into files")
	flag.BoolVar(&options.ignoreCase, "ignore-case", false, "ignore case we searching into files")

	// if set, no case sensitive
	flag.BoolVar(&options.invertMatch, "v", false, "invert the sense of matching")
	flag.BoolVar(&options.invertMatch, "invert-match", false, "invert the sense of matching")

	// if set, we only want number of matches
	flag.BoolVar(&options.matchesNumber, "c", false, "number of matches")
	flag.BoolVar(&options.matchesNumber, "count", false, "number of matches")

	// if set, we only to display the file names
	flag.BoolVar(&options.onlyFiles, "l", false, "only print files machting the regex")
	flag.BoolVar(&options.onlyFiles, "files-with-matches", false, "only print files machting the regex")

	// if set, we only to display the file names
	flag.BoolVar(&options.onlyMatching, "o", false, "only print matching tokens")
	flag.BoolVar(&options.onlyMatching, "only-matching", false, "only print matching tokens")

	flag.Usage = func() {
		fmt.Print(Usage)
	}

	flag.Parse()

	// nothing entered: show help
	if len(flag.Args()) == 0 {
		fmt.Print(Usage)
		os.Exit(0)
	}

	// manage first non-flag argument: the regexp
	re := flag.Arg(0)

	// add ignore case (?i) if specified
	if options.ignoreCase {
		options.re = regexp.MustCompile("(?i)" + re)
	} else {
		options.re = regexp.MustCompile(re)
	}
	
	options.Files = flag.Args()[1:]

	return options
}


