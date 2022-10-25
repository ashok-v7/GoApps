package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// represents a Trade
type Trade struct {
	Symbol string
	Volume int
	Price  float64
}

func genReport(w io.Writer, trades []Trade) {
	for i, t := range trades {

		fmt.Println(i, t)
		log.Printf("%d: %#v", i, t)

		fmt.Fprintf(w, "%-10s %04d %.2f\n", t.Symbol, t.Volume, t.Price)
	}
}

func main() {
	log.SetPrefix("LOG:")
	trades := []Trade{{"MSFT", 231, 234.56}, {"TSLA", 124, 345.62}, {"BRK", 65, 85.86}} // strucut initialization
	//fmt.Println(trades)
	genReport(os.Stdout, trades)
}
