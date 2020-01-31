package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func hasOption(args []string, option string) bool {
	for _, arg := range args {
		if arg == option {
			return true
		}
	}
	return false
}

func timestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func main() {
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// Input is being piped in
		args := os.Args[1:]
		var file *bufio.Writer
		if len(args) > 0 {
			fileName := args[0]
			f, _ := os.Create(fileName)
			file = bufio.NewWriter(f)
			if file != nil {
				defer file.Flush()
			}
		}
		timestamps := hasOption(args, "-t")
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			text := reader.Text()
			if timestamps {
				text = timestamp() + " - " + text
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
