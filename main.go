package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	version := "dmngo 0.0.1"

	tld := flag.String("tld", "", "Top-level domain (e.g., .com, .net, .org)")
	length := flag.Int("length", 0, "Try generated strings of this length")
	input := flag.String("input", "", "Read strings from this file.")
	vowelEnding := flag.Bool("vowel", false, "Ensure the string ends with a vowel. Possibly useful for random strings.")
	versionFlag := flag.Bool("version", false, "Print version")

	flag.Parse()

	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	// add dot before tld if it's missing
	if !strings.HasPrefix(*tld, ".") {
		*tld = "." + *tld
	}

	if *length > 0 {
		// random string lookup
		listToTry := generateCombinations(*length)

		for i := range listToTry {
			if *vowelEnding {
				vowels := "aeiou"
				// try every vowel
				for v := range vowels {
					// is the last character of the string 'v'?
					if string(listToTry[i][len(listToTry[i])-1]) == string(vowels[v]) {
						checkDomain(listToTry[i]+*tld, 5)
					}
				}
			} else {
				// no vowel check, just go
				checkDomain(listToTry[i]+*tld, 5)
			}
		}
	}

	if *input != "" {
		lines := parseFile(*input)
		for l := range lines {
			checkDomain(string(lines[l])+*tld, 5)
		}
	}
	os.Exit(0)
}
