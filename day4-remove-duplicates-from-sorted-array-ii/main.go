package main

import "fmt"

func removeDuplicates(nums []int) int {
	check := 0
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[check] == nums[i] {
			count++
			if count <= 1 {
				check++
				nums[check] = nums[i]
			}
		} else {
			count = 0
			check++
			nums[check] = nums[i]
		}
	}
	return check + 1
}

func main() {

	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}
