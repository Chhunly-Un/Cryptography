package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/sha3"
	"encoding/hex"
	"fmt"
)

func main() {
	var input1, input2 string

	fmt.Println("======== Name + Hashing Program ========")
	fmt.Print("Please input value 1: ")
	fmt.Scanln(&input1)
	fmt.Print("Please input value 2: ")
	fmt.Scanln(&input2)

	// Convert inputs to bytes
	data1 := []byte(input1)
	data2 := []byte(input2)

	// Compute all hashes
	hashMD5_A := md5.Sum(data1)
	hashMD5_B := md5.Sum(data2)

	hashSHA1_A := sha1.Sum(data1)
	hashSHA1_B := sha1.Sum(data2)

	hashSHA256_A := sha256.Sum256(data1)
	hashSHA256_B := sha256.Sum256(data2)

	hashSHA512_A := sha512.Sum512(data1)
	hashSHA512_B := sha512.Sum512(data2)

	hashSHA3_A := sha3.Sum256(data1)
	hashSHA3_B := sha3.Sum256(data2)

	// Display results in your format
	fmt.Println()
	displayHash("MD5", hashToString(hashMD5_A[:]), hashToString(hashMD5_B[:]))
	displayHash("SHA1", hashToString(hashSHA1_A[:]), hashToString(hashSHA1_B[:]))
	displayHash("SHA256", hashToString(hashSHA256_A[:]), hashToString(hashSHA256_B[:]))
	displayHash("SHA512", hashToString(hashSHA512_A[:]), hashToString(hashSHA512_B[:]))
	displayHash("SHA3-256", hashToString(hashSHA3_A[:]), hashToString(hashSHA3_B[:]))
}

// Converts hash bytes to hex string
func hashToString(hash []byte) string {
	return hex.EncodeToString(hash)
}

// Displays formatted hash output
func displayHash(name, a, b string) {
	fmt.Printf("Hash (%s):\n", name)
	fmt.Printf("Output A = %s\n", a)
	fmt.Printf("Output B = %s\n", b)

	if a == b {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}
	fmt.Println()
}
