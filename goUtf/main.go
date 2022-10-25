package main

import (
	"fmt"
	"unicode/utf8"
)

func lineLength(words []string) int {
	total := 0
	for _, word := range words {
		total += utf8.RuneCountInString(word)
		fmt.Println(total)
	}

	fmt.Println("words", len(words))

	numSpaces := len(words) - 1
	return total + numSpaces
}
func main() {
	words := []string{"Hello"} // slice
	fmt.Println(lineLength(words))

}
