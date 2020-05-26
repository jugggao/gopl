package main

import (
	"bufio"
	"fmt"
	"os"
)

type line struct {
	FileName string
	String   string
}

func main() {
	counts := make(map[line]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "ARGS")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, f.Name())
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[line]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[line{
			FileName: filename,
			String:   input.Text(),
		}]++
	}
}
