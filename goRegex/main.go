package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

var TransRe = regexp.MustCompile(`(\d+) shares of ([A-Z]+) for \$(\d+(\.\d+)?)`) // rEGUALR EXPRESSION  USAGE

type Transaction struct { // STRUCT DEFINE
	Symbol string
	Volume int
	Price  float64
}

func parseLine(line string) (Transaction, error) { // FUNC

	matches := TransRe.FindStringSubmatch(line)
	if matches == nil {
		return Transaction{}, fmt.Errorf("bad line: %q", line)
	}

	fmt.Println(matches[1])

	fmt.Println(matches[2])
	fmt.Println(matches[3])

	var t Transaction

	t.Volume, _ = strconv.Atoi(matches[1])
	t.Symbol = matches[2]
	t.Price, _ = strconv.ParseFloat(matches[3], 64)
	return t, nil
}

func main() {

	line := "12 shares of MSFT for $234.57"

	t, err := parseLine(line)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)

	fmt.Printf("%+v\n", t)

}
