package main

import "fmt"

func twoSum(nums []int, target int) []int {
	data := make(map[int]int)

	for index := 0; index < len(nums); index++ {
		remainder := target - nums[index]
		if val, ok := data[remainder]; ok && index != val {
			return []int{val, index}
		}

		data[nums[index]] = index
	}

	return []int{}
}

func main() {
	fmt.Printf("%v", twoSum([]int{0, 2, 1, 15}, 17))
}
