package main

import "fmt"

func majorityElement(nums []int) int {
	count := 0
	check := 0

	for _, num := range nums {
		if count == 0 {
			check = num
		}

		if num == check {
			count++
		} else {
			count--
		}
	}

	return check
}

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
