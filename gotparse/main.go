package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Prasing time")

	tsmp := "June 18, 1942"
	tt, err := time.Parse("January 02, 2006", tsmp)

	if err != nil {
		fmt.Printf("error : %s\n", err)

	} else {
		fmt.Println(tt)
	}

	ds := "2700ms"
	d, err := time.ParseDuration(ds)
	if err != nil {
		fmt.Printf("error :%s\n ", err)
	} else {
		fmt.Println(d)
	}

}
