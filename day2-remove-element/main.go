package main
import "fmt"

func removeElement(nums []int, val int) int {
	l := 0
	r := len(nums) - 1
    for l <= r {
		if val == nums[l]  {
			 nums[l] = nums[r]
			 r--
		} else {
			l++
		}
	}
	return l
}

func main() {

	fmt.Println(removeElement([]int{3,2,2,3}, 3))
}