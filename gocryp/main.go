package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func mdHash(input string) string {
	byteInp := []byte(input)
	md5data := md5.Sum(byteInp)

	fmt.Println(md5data) // byte data
	return hex.EncodeToString((md5data[:]))

}
func main() {

	fmt.Println(mdHash("prescriotion"))
}
