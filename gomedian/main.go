package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Median calcualtion")

	fmt.Println(median([]float64{2, 4, 5, 6, 7, 8, 9}))
	fmt.Println(median([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9}))

}

func median(nums []float64) float64 {
	vals := make([]float64, len(nums))
	copy(vals, nums)
	sort.Float64s(vals)

	i := len(vals) / 2 // get the quotoent

	if len(vals)%2 == 1 { // get the remainder

		return vals[i]

	}

	return (vals[i-1] + vals[i]) / 2

}
