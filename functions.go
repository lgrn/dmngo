package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

// generateCombinations() returns a shuffled []string of given length
// containing all combinations of a-z lowercase.
func generateCombinations(length int) []string {
	// Define the alphabet of lowercase letters
	letters := "abcdefghijklmnopqrstuvwxyz"

	// Slice to store the generated combinations
	var combinations []string

	// Recursive function to generate combinations
	var generate func(current string, depth int)
	generate = func(current string, depth int) {
		// Base case: if depth is 0, add the current combination to the slice
		if depth == 0 {
			combinations = append(combinations, current)
			return
		}
		// Loop through each letter and recursively generate combinations
		for _, letter := range letters {
			generate(current+string(letter), depth-1)
		}
	}

	// Start generating combinations
	generate("", length)

	// Shuffle the combinations using a random source
	//randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(combinations), func(i, j int) {
		combinations[i], combinations[j] = combinations[j], combinations[i]
	})

	return combinations
}

func checkDomain(domain string, sleep int) bool {
	result, err := whois.Whois(domain)
	if err == nil {
		_, err := whoisparser.Parse(result)
		// no error from whoisparser likely means it's taken
		if err == nil {
			fmt.Println("..." + domain + " is likely taken.")
			time.Sleep(5 * time.Second)
			return false
		} else {
			fmt.Println(domain, "MAY be free. adding to domains.txt")
			file, _ := os.OpenFile("domains.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			file.WriteString(domain + "\n")
			file.Close()
			time.Sleep(time.Duration(sleep) * time.Second)
			return true
		}

	} else {
		// error from whois might mean rate limiting. wait longer.
		fmt.Println("WHOIS ERROR. WAITING 3X AS LONG: ", err)
		time.Sleep(time.Duration(sleep) * 3 * time.Second)
		return false
	}
}

func parseFile(input string) []string {
	filePath := input
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	var retval []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		retval = append(retval, scanner.Text())
	}

	return retval
}
