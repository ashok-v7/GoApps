package main

import (
	"fmt"
)

func main() {

	fmt.Println("mean calculation")
	//fmt.Println	fmt.Println((mean[]))

	fmt.Println(mean([]int{1, 2, 3, 4, 5}), "MEean")
	fmt.Println(mean([]int{4, 5, 6}))
}

func mean(nums []int) float64 {
	s := sum(nums)
	fmt.Println(s)
	return float64(s) / float64(len(nums))

}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	return total
}
