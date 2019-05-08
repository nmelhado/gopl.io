// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if counts[line] == nil {
				counts[line] = make(map[string]int)
			}
			counts[line][filename]++
		}
	}
	for line, occurences := range counts {
		if len(occurences) > 1 {
			var files []string
			for file := range occurences {
				files = append(files, file)
			}
			fmt.Printf("\"%s\" occurs multiple times accross:\n\t%v\n", line, strings.Join(files, ", "))
		} else {
			for file, occurence := range occurences {
				if occurence > 1 {
					fmt.Printf("\"%s\" occurs multiple times accross:\n\t%v\n", line, file)
				}
			}
		}
	}
}

//!-
