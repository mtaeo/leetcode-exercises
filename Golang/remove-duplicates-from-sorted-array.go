package main

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...) // re-slicing i.e. shifting to the left
			i--
		}
	}
	return len(nums)
}

func main() {
	print(removeDuplicates(([]int{1, 1})))
}
