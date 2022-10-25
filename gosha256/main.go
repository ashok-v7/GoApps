package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func sha256hash(inp string) string {
	plainText := []byte(inp) //  input is converted to a slice of bytes

	sha256data := sha256.Sum256(plainText)

	fmt.Println(sha256data)

	return hex.EncodeToString(sha256data[:])

}
func main() {

	fmt.Println(sha256hash("treatment details"))

}
