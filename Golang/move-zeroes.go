package main

func moveZeroes(nums []int) []int {
	lastNonZeroFoundAt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[lastNonZeroFoundAt] = nums[lastNonZeroFoundAt], nums[i]
			lastNonZeroFoundAt++
		}
	}

	return nums
}

func main() {
	print(moveZeroes([]int{0, 1, 0, 3, 12}))
}
