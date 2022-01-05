package main

func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums)

	for left < right {
		middle := left + (right-left)/2

		if nums[middle] == target {
			return middle
		}

		if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return left
}

func main() {
	print(searchInsert([]int{1, 3, 5, 6}, 2))
}
