package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"password-cracker/utils/crack"
)

const (
	targetSHA1 = "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	wordlist   = "nord_vpn.txt" // put the downloaded file here
)

var trialCount int // for verbose numbering

func main() {
	start := time.Now()

	f, err := os.Open(wordlist)
	if err != nil {
		log.Fatalf("cannot open wordlist: %v", err)
	}
	defer f.Close()

	// The scanner also opens verbose.txt
	scanner := crack.NewLineScanner(f)
	defer scanner.Close()

	fmt.Printf("Cracking SHA-1 %s …\n", targetSHA1)

	password, found := crack.CrackSHA1(scanner, targetSHA1)
	duration := time.Since(start)

	if found {
		fmt.Printf("\n*** PASSWORD FOUND ***\n%s\n", password)
		fmt.Printf("Tried %d passwords in %v\n", trialCount, duration)
	} else {
		fmt.Printf("\nPassword not in wordlist (tried %d entries, %v)\n", trialCount, duration)
	}
}