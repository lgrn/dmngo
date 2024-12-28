package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

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

func checkDomain(domain string, sleep int, debugFlag **bool) bool {
	whoisResult, err := whois.Whois(domain)
	if err == nil {
		parseResult, err := whoisparser.Parse(whoisResult)
		// no error from whoisparser likely means it's taken
		if err == nil {
			//fmt.Println("..." + domain + " is likely taken.")
			fmt.Printf("TAKEN: %s\tuntil: %s\n",
				domain,
				//parseResult.Domain.Status,
				//registrantName,
				parseResult.Domain.ExpirationDate)
			if **debugFlag {
				// Print the domain status
				fmt.Println(parseResult.Domain.Status)

				// Print the domain created date
				fmt.Println(parseResult.Domain.CreatedDate)

				// Print the domain expiration date
				fmt.Println(parseResult.Domain.ExpirationDate)

				// Print the registrar name
				fmt.Println(parseResult.Registrar.Name)

				// Print the registrant name
				fmt.Println(parseResult.Registrant.Name)

				// Print the registrant email address
				fmt.Println(parseResult.Registrant.Email)

			}
			time.Sleep(5 * time.Second)
			return false
		} else {
			fmt.Println(domain, "MAY be free. adding to domains.txt")
			fmtError := fmt.Errorf("additional context: %w", err)
			fmt.Println(fmtError)
			file, _ := os.OpenFile("domains.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			file.WriteString(domain + "\n")
			file.Close()
			time.Sleep(time.Duration(sleep) * time.Second)
			return true
		}

	} else {
		parseResult, err := whoisparser.Parse(whoisResult)
		if err == nil {
			// parse was OK
			// we did get an error from the whois though, which might
			// mean rate limiting. wait longer.
			fmt.Println("WHOIS ERROR. WAITING 3X AS LONG")
			fmt.Printf("ERROR: %s\t status: %s\n",
				domain,
				parseResult.Domain.Status,
			//registrantName,
			//parseResult.Domain.ExpirationDate
			)
			fmt.Println(fmt.Errorf("additional context: %w", err))
		} else {
			fmt.Println("PARSE ERROR")
			fmt.Println(fmt.Errorf("additional context: %w", err))
		}
		time.Sleep(time.Duration(sleep) * 3 * time.Second)
		return false
	}
}

func toLowercaseASCII(s string) string {
	var builder strings.Builder
	for _, r := range s {
		// Convert to lowercase and ensure it's an ASCII letter
		if unicode.IsLetter(r) && r <= 'z' {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
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
		cleanLine := toLowercaseASCII(scanner.Text())
		retval = append(retval, cleanLine)
	}

	rand.Shuffle(len(retval), func(i, j int) {
		retval[i], retval[j] = retval[j], retval[i]
	})

	return retval
}
