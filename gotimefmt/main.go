package main

import (
	"fmt"
	"time"
)

func main() {
	renon := time.Date(2001, time.September, 11, 11, 0, 0, 0, time.UTC)
	fmt.Println(renon) // 1974-09-11 11:00:00 +0000 UTC

	fmt.Println(renon.Format("2022-01-02"))
	fmt.Println(renon.Format("Sun,Oct 23"))

	fmt.Println(renon.Format(time.RFC3339Nano))

	dg := 3500 * time.Millisecond

	fmt.Println(dg)

}
