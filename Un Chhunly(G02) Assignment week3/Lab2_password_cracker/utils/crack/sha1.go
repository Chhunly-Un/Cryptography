package crack

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// CrackSHA1 tries every line from r (a wordlist) until the SHA-1 matches targetHash.
// It returns the matching password and a boolean indicating success.
func CrackSHA1(r io.Reader, targetHash string) (string, bool) {
	// Decode the target once
	targetBytes, _ := hex.DecodeString(targetHash)

	scanner := NewLineScanner(r)
	line := 0
	for scanner.Scan() {
		line++
		password := scanner.Text()
		h := sha1.Sum([]byte(password))
		if hex.EncodeToString(h[:]) == targetHash {
			return password, true
		}
	}
	return "", false
}

// Helper – a tiny line scanner that also prints verbose info to a file.
type lineScanner struct {
	*bufio.Scanner
	file *os.File
}

func NewLineScanner(r io.Reader) *lineScanner {
	f, _ := os.OpenFile("verbose.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	return &lineScanner{
		Scanner: bufio.NewScanner(r),
		file:    f,
	}
}

func (ls *lineScanner) Scan() bool {
	ok := ls.Scanner.Scan()
	if ok {
		line := ls.Text()
		// verbose line:  trial#  password
		io.WriteString(ls.file, fmt.Sprintf("Trial %d: %s\n", trialCount, line))
		trialCount++
	}
	return ok
}

// close the verbose file when done
func (ls *lineScanner) Close() { ls.file.Close() }
