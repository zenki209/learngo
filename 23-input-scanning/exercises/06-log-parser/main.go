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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//	You've watched the lecture. Now, try to create the same
//	log parser program on your own. Do not look at the lecture,
//	and the existing source code.
//
// EXPECTED OUTPUT
//
//	go run main.go < log.txt
//
//	 DOMAIN                             VISITS
//	 ---------------------------------------------
//	 blog.golang.org                        30
//	 golang.org                             10
//	 learngoprogramming.com                 20
//
//	 TOTAL                                  60
//
//
//	go run main.go < log_err_missing.txt
//
//	 wrong input: [golang.org] (line #3)
//
//
//	go run main.go < log_err_negative.txt
//
//	 wrong input: "-100" (line #3)
//
//
//	go run main.go < log_err_str.txt
//
//	 wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------
type parse struct {
	DOMAINS map[string]int
	domain  []string
	total   int
	line    int
}

func main() {

	data := bufio.NewScanner(os.Stdin)

	p := parse{
		DOMAINS: make(map[string]int),
	}

	line := 0
	for data.Scan() {
		line++
		parts := strings.Split(data.Text(), " ")
		if len(parts) != 2 {
			fmt.Printf("wrong input: %q (line #%d)\n", data.Text(), line)
			continue
		}
		domain := parts[0]
		visits, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", parts[1], line)
			continue
		}
		p.DOMAINS[domain] += visits
		p.domain = append(p.domain, domain)
		p.total += visits
	}
	fmt.Printf("%-30s %s\n", "DOMAIN", "VISITS")
	fmt.Println("---------------------------------------------")

	for _, domain := range p.domain {
		fmt.Printf("%-30s %d\n", domain, p.DOMAINS[domain])
	}
	fmt.Println("---------------------------------------------")
	fmt.Printf("%-30s %d\n", "TOTAL", p.total)
}
