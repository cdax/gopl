package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, filemap := range counts {
		filewise_counts := make([]string, len(filemap))
		i, sum := 0, 0
		for file, n := range filemap {
			filewise_counts[i] = fmt.Sprintf("%s(%d)", file, n)
			i++
			sum += n
		}
		if sum > 1 {
			fmt.Printf("%s\t%s\n", line, strings.Join(filewise_counts, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	f_name := f.Name()
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][f_name]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
