package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := reader.Text()
		fmt.Println(text)
		if file != nil {
			_, _ = file.WriteString(text + "\n")
			go file.Flush()
		}
	}
}
