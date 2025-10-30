package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"lab2_password_cracker/utils/crack"
)

func main() {
	wordlistPath := flag.String("wordlist", "nord_vpn.txt", "path to the wordlist file")
	targetHash := flag.String("hash", "", "target SHA1 hash to crack (hex)")
	outFile := flag.String("out", "verbose.txt", "file to write verbose output")
	limit := flag.Int("limit", 0, "optional: limit number of words read (0 = no limit)")
	flag.Parse()

	if *targetHash == "" {
		fmt.Print("Enter target SHA1 (hex): ")
		in := bufio.NewReader(os.Stdin)
		text, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}
		*targetHash = strings.TrimSpace(text)
	}

	*targetHash = strings.ToLower(strings.ReplaceAll(*targetHash, " ", ""))
	*targetHash = strings.ReplaceAll(*targetHash, "\n", "")
	*targetHash = strings.ReplaceAll(*targetHash, "\r", "")

	if len(*targetHash) != 40 {
		fmt.Fprintf(os.Stderr, "Warning: provided hash has length %d (expected 40 for SHA1). The program will continue.\n", len(*targetHash))
	}

	f, err := os.Open(*wordlistPath)
	if err != nil {
		log.Fatalf("Failed to open wordlist '%s': %v", *wordlistPath, err)
	}
	defer f.Close()

	verbF, err := os.Create(*outFile)
	if err != nil {
		log.Fatalf("Failed to create verbose file '%s': %v", *outFile, err)
	}
	defer verbF.Close()

	writeVerbose := func(s string) {
		fmt.Fprintln(verbF, s)
		fmt.Println(s) 
	}

	scanner := bufio.NewScanner(f)
	lineNo := 0
	found := false
	var foundPlain string
	start := time.Now()
	writeVerbose(fmt.Sprintf("Start cracking at %s", start.Format(time.RFC3339)))
	writeVerbose(fmt.Sprintf("Target SHA1: %s", *targetHash))
	writeVerbose(fmt.Sprintf("Wordlist: %s", *wordlistPath))
	writeVerbose("---- Verbose output (each checked word and its SHA1) ----")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lineNo++
		if *limit > 0 && lineNo > *limit {
			break
		}

		computed := crack.SHA1Hex(line)
		writeVerbose(fmt.Sprintf("[%d] word='%s' sha1=%s", lineNo, line, computed))

		if computed == *targetHash {
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
		fmt.Printf("\nCracked! SHA1 %s => password: %s\n", *targetHash, foundPlain)
	} else {
		fmt.Printf("\nNot found in wordlist (searched %d lines)\n", lineNo)
	}
}
