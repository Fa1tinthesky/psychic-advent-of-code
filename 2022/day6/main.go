package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"strings"
)

func main() {
	var filePath string;
	flag.StringVar(&filePath, "path", "./2022/day6/tests/example.txt", "Path to test file")
	flag.Parse()

	d, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal("Couldn't read the file")
	}

	cnt := strings.Trim(string(d), "\n")
	fmt.Println(cnt)

	fmt.Println("FOUND ON INDEX:", findSubstr(cnt))
}


func findSubstr(input string) int {
	// r := strings.NewReader(input)

	var start, end int;

	for start = 0; start < len(input) - 3; start++ {
		end = start + 3
		sub_s := input[start:end+1]

		if match(sub_s) {
			return end + 1
		}

		// add(*found, input[i])

		fmt.Println(sub_s)
	}

	return 0
}

// Goal: see all the letter in arr are different
func match(s string) bool {
	m := make(map[rune]int)

	for _, v := range s {
		if m[v] == 0 {
			m[v] = 1
		} else {
			return false
		}
	}

	return true
}

// TODO:
// Find first substring in data that has 4 different digits
// Output its last digit number in order (index + 1)

