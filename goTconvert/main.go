package main

import (
	"fmt"
	"time"
)

func tsConver(ts, from, to string) (string, error) {
	fromTz, err := time.LoadLocation(from)
	fmt.Println(fromTz)
	if err != nil {
		return "", err
	}

	toTz, err := time.LoadLocation(to)
	if err != nil {
		return "", err
	}

	const format = "2006-01-02T15:04"
	fromTime, err := time.ParseInLocation(format, ts, fromTz)
	fmt.Println(fromTime)
	if err != nil {
		return "", err
	}

	toTime := fromTime.In(toTz)
	return toTime.Format(format), nil
}

func main() {
	fmt.Println("TIME CONVERT")
	ts := "2021-03-08T19:12"

	out, err := tsConver(ts, "America/Los_Angeles", "Asia/Jerusalem")
	if err != nil {
		fmt.Printf("error : %s", err)

		return
	}
	fmt.Println(out)
}
