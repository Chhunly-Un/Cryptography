package main

import (
	"flag"
	"log"
	"os"

	"lab1_md5_cracker/utils/crack"
)

func main() {
	wordlist := flag.String("wordlist", "nord_vpn.txt", "path to wordlist")
	output := flag.String("output", "verbose.txt", "verbose output file")
	hash := flag.String("hash", "6a85dfd77d9cb35770c9dc6728d73d3f", "MD5 hash to crack")
	flag.Parse()

	f, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	found, err := crack.CrackMD5(*hash, *wordlist, f)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Password found: %s", found)
}