package crack

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

// lineScanner wraps bufio.Scanner to provide trimmed text
type lineScanner struct {
	*bufio.Scanner
}

func NewLineScanner(r io.Reader) lineScanner {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	return lineScanner{s}
}

func (ls lineScanner) Text() string {
	return strings.TrimSpace(ls.Scanner.Text())
}

// CrackMD5 tries to find the password for the given MD5 hash
func CrackMD5(hash string, wordlistPath string, verbose io.Writer) (string, error) {
	file, err := os.Open(wordlistPath)
	if err != nil {
		return "", fmt.Errorf("open wordlist: %w", err)
	}
	defer file.Close()

	scanner := NewLineScanner(file)
	lineNo := 0
	fmt.Fprintf(verbose, "Target MD5: %s\n", hash)
	fmt.Fprintf(verbose, "Wordlist : %s\n", wordlistPath)
	fmt.Fprintln(verbose, "--------------------------------------------------")

	for scanner.Scan() {
		lineNo++
		pwd := scanner.Text()
		if pwd == "" {
			continue
		}

		attempt := md5.Sum([]byte(pwd))
		attemptHex := hex.EncodeToString(attempt[:])

		fmt.Fprintf(verbose, "[%6d] trying: %-30s -> %s\n", lineNo, pwd, attemptHex)

		if attemptHex == hash {
			fmt.Fprintf(verbose, "\n*** PASSWORD FOUND ***\n")
			fmt.Fprintf(verbose, "Line: %d\nPassword: %s\nHash: %s\n", lineNo, pwd, attemptHex)
			return pwd, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("reading wordlist: %w", err)
	}
	return "", fmt.Errorf("password not found in wordlist")
}