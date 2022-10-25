package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// {
//     "time": "2020-03-06T00:00:00Z",
//     "station": "DS9",
//     "temperature": 21.6,
//     "rain": 0
//   }
type Record struct {
	Time    time.Time
	Station string
	Temp    float64 `json:"temperature"`
	Rain    float64
}

func readRecord(r io.Reader) (Record, error) {
	var rec Record
	dec := json.NewDecoder(r)
	if err := dec.Decode(&rec); err != nil {
		return Record{}, err
	}
	return rec, nil
}

func main() {

	file, err := os.Open("record.json")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	rec, err := readRecord(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rec)

}
