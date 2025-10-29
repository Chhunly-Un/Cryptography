package main

import (
    "encoding/hex"
    "fmt"
    "net/url"
)

func main() {
    // Percent-encoded string (example)
    encoded := "%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32"

    // Step 1: URL decode
    decodedURL, err := url.QueryUnescape(encoded)
    if err != nil {
        panic(err)
    }
    fmt.Println("URL-decoded string:", decodedURL)

    // Step 2: Use regex hint \x6d\x65\x6f\x77 -> hex "6d656f77"
    hexString := "6d656f77" // from regex hint
    bytes, err := hex.DecodeString(hexString)
    if err != nil {
        panic(err)
    }

    // Step 3: Repeat twice
    repeated := string(bytes) + string(bytes)

    // Step 4: Format flag
    flag := fmt.Sprintf("cryptoCTF{%s}", repeated)
    fmt.Println("Flag:", flag)
}
