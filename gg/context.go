package gg

import (
	"fmt"
	"strings"
)

// ASCII terminal colors
const (
	Black   = "\033[1;30m%s\033[0m"
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	Magenta = "\033[1;35m%s\033[0m"
	Teal    = "\033[1;36m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
)

// this will hold the current context when searching
type Context struct {
	// current file name being search into
	fileName *string

	// current line
	currentLine string

	// current line number
	lineNumber int64

	// current number of matches
	matches int64
}

// search into the open file
func (ctx *Context) Grep(options *CliOptions) {
	if ctx.SimpleMatch(options) {
		ctx.matches++
		ctx.DisplayInfo(options)
	}
}

// simple match which can be inverted
func (ctx *Context) SimpleMatch(options *CliOptions) bool {
	if options.invertMatch {
		return !options.re.MatchString(ctx.currentLine)
	} else {
		return options.re.MatchString(ctx.currentLine)
	}
}

func (ctx *Context) DisplayInfo(options *CliOptions) {
	// number of matches: we'll display output at the end of the file search
	if options.matchesNumber || options.onlyFiles {
		return
	}

	// get any capturing group value to display it with colors
	matches := options.re.FindAllString(ctx.currentLine,-1)
	//fmt.Printf("matches=%#v", options.re.FindAllString(ctx.currentLine,-1))
	for _,m := range matches {
		found_string := fmt.Sprintf(Yellow, m)
		ctx.currentLine = strings.Replace(ctx.currentLine, m, found_string, -1)
	}

	if options.requestLinesNumbers {
		fmt.Printf("%s(%d):%s\n", *ctx.fileName, ctx.lineNumber, ctx.currentLine)
	} else {
		format := fmt.Sprintf("%s:%%s\n", Green)
		fmt.Printf(format, *ctx.fileName, ctx.currentLine)
	}
}

// display file name and number of matches only
func (ctx *Context) DisplayMatches() {
	yellowFormat := fmt.Sprintf(Yellow, "%d")
	format := fmt.Sprintf("%s:%s\n", Green, yellowFormat)
	fmt.Printf(format, *ctx.fileName, ctx.matches)
}

// display file names
func (ctx *Context) DisplayFileNames() {
	format := fmt.Sprintf("%s\n", Green)
	fmt.Printf(format, *ctx.fileName)
}
