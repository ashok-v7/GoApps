package main

import (
	"fmt"
	"strings"
)

type Letter struct {
	Symbol  string
	English string
}

var letters = []Letter{{"Σ", "Sigma"}}

func englishFor(greek string) (string, error) {
	for _, letter := range letters {
		if strings.EqualFold(greek, letter.Symbol) {
			return letter.English, nil
		}
	}
	return "", fmt.Errorf("unknown geek letter : %#v", greek)
}
func main() {

	fmt.Println(englishFor("Σ"))
	fmt.Println(englishFor("σ"))
	fmt.Println(englishFor("ς"))

}
