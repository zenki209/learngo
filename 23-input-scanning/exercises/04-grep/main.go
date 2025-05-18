// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)

	// Collect all command-line arguments as patterns
	patterns := os.Args[1:]

	for in.Scan() {
		s := in.Text()
		// Check if any of the patterns are present in the line
		for _, pattern := range patterns {
			if strings.Contains(s, pattern) {
				fmt.Printf("Found the word \"%s\" in sentence: %s\n", pattern, s)
				break // Stop checking other patterns if one is found
			}
		}
	}
}
