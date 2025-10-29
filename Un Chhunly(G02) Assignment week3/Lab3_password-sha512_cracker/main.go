package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"Lab3_password-sha512_cracker/utils/crack"
)

func main() {
	wordlistPath := flag.String("wordlist", "nord_vpn.txt", "path to the wordlist file")
	targetHash := flag.String("hash", "", "target SHA512 hash to crack (hex)")
	outFile := flag.String("out", "verbose.txt", "file to write verbose output")
	limit := flag.Int("limit", 0, "optional: limit number of words read (0 = no limit)")
	flag.Parse()

	// Prompt for SHA512 if not provided via flag
	if *targetHash == "" {
		fmt.Print("Enter target SHA512 (hex): ")
		in := bufio.NewReader(os.Stdin)
		text, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}
		*targetHash = strings.TrimSpace(text)
	}

	// normalize target: remove spaces/newlines and lowercase
	normalized := strings.ToLower(strings.ReplaceAll(*targetHash, " ", ""))
	normalized = strings.ReplaceAll(normalized, "\n", "")
	normalized = strings.ReplaceAll(normalized, "\r", "")

	// Optional sanity check for SHA512 length (128 hex chars)
	if len(normalized) != 128 {
		fmt.Fprintf(os.Stderr, "Warning: provided hash length = %d (expected 128 for SHA-512). Continuing anyway.\n", len(normalized))
	}

	// Open wordlist
	f, err := os.Open(*wordlistPath)
	if err != nil {
		log.Fatalf("Failed to open wordlist '%s': %v", *wordlistPath, err)
	}
	defer f.Close()

	// Prepare verbose output file
	verbF, err := os.Create(*outFile)
	if err != nil {
		log.Fatalf("Failed to create verbose file '%s': %v", *outFile, err)
	}
	defer verbF.Close()

	writeVerbose := func(s string) {
		fmt.Fprintln(verbF, s)
		fmt.Println(s) // print also to stdout
	}

	scanner := bufio.NewScanner(f)
	lineNo := 0
	found := false
	var foundPlain string
	start := time.Now()
	writeVerbose(fmt.Sprintf("Start cracking at %s", start.Format(time.RFC3339)))
	writeVerbose(fmt.Sprintf("Target SHA512: %s", normalized))
	writeVerbose(fmt.Sprintf("Wordlist: %s", *wordlistPath))
	writeVerbose("---- Verbose output (each checked word and its SHA512) ----")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lineNo++
		if *limit > 0 && lineNo > *limit {
			break
		}

		computed := crack.SHA512Hex(line)
		writeVerbose(fmt.Sprintf("[%d] word='%s' sha512=%s", lineNo, line, computed))

		if computed == normalized {
			found = true
			foundPlain = line
			writeVerbose("---- MATCH FOUND ----")
			writeVerbose(fmt.Sprintf("Line %d: %s -> %s", lineNo, line, computed))
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning wordlist: %v", err)
	}

	elapsed := time.Since(start)
	writeVerbose(fmt.Sprintf("Finished at %s (elapsed %s)", time.Now().Format(time.RFC3339), elapsed))

	if found {
		fmt.Printf("\nCracked! SHA512 %s => password: %s\n", normalized, foundPlain)
	} else {
		fmt.Printf("\nNot found in wordlist (searched %d lines)\n", lineNo)
	}
}
