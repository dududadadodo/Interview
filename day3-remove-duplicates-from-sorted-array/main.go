package main
import "fmt"

func removeDuplicates(nums []int) int {
	check := 1
    for i := 0; i < len(nums); i++ {
      if nums[check-1] != nums[i]  {
		nums[check] = nums[i]
		check++
      } 
    }
	return check
}

func main() {

	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}