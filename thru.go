package main

import (
	"bufio"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/fatih/color"
	"github.com/kamackay/thru/model"
	"github.com/kamackay/thru/version"
	"os"
	"regexp"
	"strings"
	"time"
)

func timestamp() string {
	now := time.Now().UTC()
	return fmt.Sprintf(now.Format("2006-02-01 15:04:05.000"))
}

func main() {
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	var opts model.Opts
	_ = kong.Parse(&opts)
	if opts.Version {
		fmt.Printf("%s\n", version.VERSION)
		return
	}
	fileName := opts.File
	fi, _ := os.Stdin.Stat()

	var (
		enableHighlight                = len(opts.Highlight) > 0
		re              *regexp.Regexp = nil
		err             error
		highlightRegex  = true
	)
	if enableHighlight {
		re, err = regexp.Compile(opts.Highlight)
		if err != nil {
			highlightRegex = true
		}
	}
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// Input is being piped in
		var file *bufio.Writer
		if len(fileName) > 0 {
			f, _ := os.Create(fileName)
			file = bufio.NewWriter(f)
			if file != nil {
				defer file.Flush()
			}
		}
		timestamps := opts.Timestamps
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			text := reader.Text()
			if enableHighlight {
				if highlightRegex {
					text = re.ReplaceAllStringFunc(text, func(s string) string {
						return yellow.Sprintf(s)
					})
				} else {
					text = strings.ReplaceAll(text, opts.Highlight, red.Sprintf(opts.Highlight))
				}
			}
			if timestamps {
				text = fmt.Sprintf("%s - %s", timestamp(), text)
			}
			fmt.Println(text)
			if file != nil {
				_, _ = file.WriteString(text + "\n")
				file.Flush()
			}
		}
	} else {
		// Input is not coming in through Pipe
		return
	}
}
